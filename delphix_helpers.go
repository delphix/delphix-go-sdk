package delphix

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	resty "gopkg.in/resty.v1"
)

// WaitforDelphixJob waits for a job to complete
func (c *Client) WaitforDelphixJob(j string) error {
	var jobState string
	var err error
	for jobState != "COMPLETED" && jobState != "FAILED" && jobState != "CANCELED" {
		time.Sleep(3 * time.Second)
		resp, err := resty.R().
			SetHeader("Content-Type", "application/json").
			Get(c.url + "/job/" + j)
		// explore response object
		if err != nil {
			panic(err)
		}
		s := resp.Body()

		var dat map[string]interface{}
		if err = json.Unmarshal(s, &dat); err != nil { //convert the json to go objects
			return err
		}
		results := dat["result"].(map[string]interface{}) //grab the query results
		jobState = results["jobState"].(string)
		fmt.Printf("\n%v", results["jobState"])
	}
	//If the job is failed or cancelled, return an error
	if jobState == "FAILED" {
		err = fmt.Errorf("Job Failed")
	} else if jobState == "CANCELED" {
		err = fmt.Errorf("Job Canceled")
	}
	return err
}

// CreateDatabase provisions an Oracle virtual database
func (c *Client) CreateDatabase(o *OracleProvisionParametersStruct) (
	interface{}, error) {
	reference, err := c.executePostJobAndReturnObjectReference("/database/provision", o)
	return reference, err
}

// UpdateDatabase updates an Oracle virtual database object
func (c *Client) UpdateDatabase(r string, o *OracleDatabaseContainerStruct) error {
	url := fmt.Sprintf("/database/%s", r)
	err := c.executePostJobAndReturnErrOnly(url, o)
	return err

}

// DeleteDatabase deletes a virtual database
func (c *Client) DeleteDatabase(v string) error {

	o := OracleDeleteParametersStruct{
		Type: "OracleDeleteParameters",
	}
	url := fmt.Sprintf("/database/%s/delete", strings.ToUpper(v))
	err := c.executePostJobAndReturnErrOnly(url, o)
	return err
}

// CheckStoragePool checks whether the storage domain has been created
func (c *Client) CheckStoragePool() (interface{}, error) {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		Get(c.url + "/domain")
	// explore response object
	if err != nil {
		panic(err)
	}
	s := resp.Body()

	var dat map[string]interface{}
	if err = json.Unmarshal(s, &dat); err != nil { //convert the json to go objects
		return nil, err
	}
	if dat["result"] != nil {
		results := dat["result"].(map[string]interface{}) //grab the query results
		fmt.Printf("Domain %s exists.\n", results["name"])
	} else {
		fmt.Printf("%s\n", dat["error"].(map[string]interface{})["details"])
	}

	return s, nil
}

// InitializeSystem completes the initial setup wizard
func (c *Client) InitializeSystem(d string, p string) (interface{}, error) {
	body, err := CreateSystemInitializationParameters(c, d, p)
	if err != nil {
		return nil, err
	}
	fmt.Println("Assigning devices to storage domain")
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(c.url + "/domain/initializeSystem")

	if err != nil {
		return nil, err
	}

	if http.StatusOK != resp.StatusCode() { //check to make sure our query was good
		errorMessage := string(resp.Body())
		err = fmt.Errorf(errorMessage)
		if err != nil {
			return nil, err
		}
	}

	result := resp.Body()

	var resultdat map[string]interface{}
	if err = json.Unmarshal(result, &resultdat); err != nil { //convert the json to go objects
		return nil, err
	}

	if resultdat["status"].(string) == "ERROR" {
		errorMessage := string(resp.Body())
		err = fmt.Errorf(errorMessage)
		if err != nil {
			return nil, err
		}
	}

	c.WaitForEngineReady(10, 300)

	return result, err
}

// FactoryReset performs a factory reset on the Delphix engine
func (c *Client) FactoryReset() error {
	fmt.Println("Restoring Delphix DDP to factory defaults")
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{}`).
		Post(c.url + "/system/factoryReset")
	if err != nil {
		return err
	}

	if http.StatusOK != resp.StatusCode() { //check to make sure our query was good
		errorMessage := string(resp.Body())
		err = fmt.Errorf(errorMessage)
		if err != nil {
			return err
		}
	}

	result := resp.Body()

	var resultdat map[string]interface{}
	if err = json.Unmarshal(result, &resultdat); err != nil { //convert the json to go objects
		return err
	}

	if resultdat["status"].(string) == "ERROR" {
		errorMessage := string(resp.Body())
		err = fmt.Errorf(errorMessage)
		if err != nil {
			return err
		}
	}

	c.WaitForEngineReady(10, 300)
	return nil
}

// WaitForEngineReady loops until the Client connection is successful or time (t) expires
func (c *Client) WaitForEngineReady(p int, t int) error {

	log.Printf("Waiting up to %v seconds for the DDDP to be ready", t)
	timeOut := 0
	for timeOut < t {
		fmt.Println("Waiting for Delphix DDP")
		time.Sleep(time.Duration(p) * time.Second)
		timeOut = timeOut + p
		if err := c.LoadAndValidate(); err == nil {
			break
		} else if err.Error() == "Delphix Username/Password incorrect" {
			return err
		}
	}
	return nil
}

// GetStorageDevices returns the list of unassigned storage devices on the Delphix engine
func (c *Client) GetStorageDevices() ([]string, error) {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		Get(c.url + "/storage/device")
	// explore response object
	if err != nil {
		panic(err)
	}
	s := resp.Body()
	//fmt.Printf("Body: %v", string(s))
	var resultdat map[string]interface{}
	if err = json.Unmarshal(s, &resultdat); err != nil { //convert the json to go objects
		return nil, err
	}

	if resultdat["status"].(string) == "ERROR" {
		errorMessage := string(resp.Body())
		err = fmt.Errorf(errorMessage)
		if err != nil {
			return nil, err
		}
	}
	results := resultdat["result"].([]interface{}) //grab the query results
	devices := []string{}
	fmt.Println("The following devices are available:")
	for _, result := range results { //loop through the results
		if result.(map[string]interface{})["configured"] != true { //if the device is not already configured
			fmt.Println(result.(map[string]interface{})["reference"]) //grab the device reference
			devices = append(devices, result.(map[string]interface{})["reference"].(string))
		}
	}
	return devices, nil
}

// CreateSystemInitializationParameters assembles the object necessary to feed the InitializeSystem function
func CreateSystemInitializationParameters(c *Client, d string, p string) (string, error) {
	devices, err := c.GetStorageDevices()
	if len(devices) == 0 {
		err = fmt.Errorf("No devices available for assignment into the domain")
	}
	if err != nil {
		return "", err
	}
	var deviceString string
	for i, device := range devices {
		device = "\"" + device + "\""
		if i == 0 {
			deviceString = device
			continue
		}
		deviceString = deviceString + "," + device
	}
	return fmt.Sprintf(`
		{
					"type": "SystemInitializationParameters",
					"defaultUser": "%s",
					"defaultPassword": "%s",
					"devices": [
						%s
					]
				}
		`, d, p, deviceString), nil
}

// UpdateUserPasswordByName updates the specified username's(u) password(p)
func (c *Client) UpdateUserPasswordByName(u string, p string) error {
	userRef, err := c.FindUserByName(u)
	if err != nil {
		return err
	}
	if userRef == nil {
		log.Fatalf("%s not found. Exiting", u)
	}
	fmt.Printf("Changing %s's password\n", u)
	fmt.Println("Using " + c.username)
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(fmt.Sprintf(`
			{
				"type": "CredentialUpdateParameters",
				"newCredential": {
					"type": "PasswordCredential",
					"password": "%s"
				}
			}
		`, p)).
		Post(c.url + "/user/" + userRef.(string) + "/updateCredential")
	if err != nil {
		return err
	}

	if http.StatusOK != resp.StatusCode() { //check to make sure our query was good
		errorMessage := string(resp.Body())
		err = fmt.Errorf(errorMessage)
		if err != nil {
			return err
		}
	}

	result := resp.Body()

	var resultdat map[string]interface{}
	if err = json.Unmarshal(result, &resultdat); err != nil { //convert the json to go objects
		return err
	}

	if resultdat["status"].(string) == "ERROR" {
		errorMessage := string(resp.Body())
		err = fmt.Errorf(errorMessage)
		if err != nil {
			return err
		}
	}
	//If we just changed the password for the current user, use that password for future connections in this session
	if c.username == u {
		c.password = p
		c.LoadAndValidate()
	}
	return nil
}

// GrabObjectNameAndReference parses a results object(o) and returns the name and reference of that object in Delphix
func GrabObjectNameAndReference(o interface{}) (string, string, error) {
	var objName, objReference interface{}
	objMap := o.(map[string]interface{})

	if objName = objMap["name"]; objName == nil {
		return "", "", fmt.Errorf("Object does not have a name")
	}
	if objReference = objMap["reference"]; objReference == nil {
		return "", "", fmt.Errorf("Object does not have a reference")
	}
	return objName.(string), objReference.(string), nil
}

// CreateEnvironment creates a new environment in Delphix
func (c *Client) CreateEnvironment(h *HostEnvironmentCreateParametersStruct) (
	interface{}, error) {
	url := "/environment"
	reference, err := c.executePostJobAndReturnObjectReference(url, h)

	return reference, err
}

// DeleteEnvironment deletes an environment in Delphix
func (c *Client) DeleteEnvironment(r string) error {
	url := fmt.Sprintf("/environment/%s/delete", strings.ToUpper(r))
	err := c.executePostJobAndReturnErrOnly(url, "{}")

	return err
}

// UpdateEnvironment updates an environment in Delphix
func (c *Client) UpdateEnvironment(r string, h *UnixHostEnvironmentStruct) error {
	url := fmt.Sprintf("/environment/%s", r)

	err := c.executePostJobAndReturnErrOnly(url, h)

	return err
}

// CreateGroup creates a new group in Delphix
func (c *Client) CreateGroup(h *GroupStruct) (
	interface{}, error) {
	url := "/group"
	reference, err := c.executePostJobAndReturnObjectReference(url, h)

	return reference, err
}

// DeleteGroup deletes an group in Delphix
func (c *Client) DeleteGroup(r string) error {
	url := fmt.Sprintf("/group/%s/delete", strings.ToUpper(r))
	err := c.executePostJobAndReturnErrOnly(url, "{}")

	return err
}

// UpdateGroup updates an group in Delphix
func (c *Client) UpdateGroup(r string, h *GroupStruct) error {
	url := fmt.Sprintf("/group/%s", r)

	err := c.executePostJobAndReturnErrOnly(url, h)

	return err
}

// CreateDSource creates a dSource in Delphix
func (c *Client) CreateDSource(l *LinkParametersStruct) (
	interface{}, error) {

	sync := false
	resultdat, err := mapifyStruct(l)
	log.Println("resultdat")
	log.Println(resultdat)
	if err != nil {
		return nil, err
	}
	//If they set "LinkNow" to true, we want to intercept that because we can't
	//catch the job and wait for it to finish, so we will link, then sync
	if resultdat["linkData"].(map[string]interface{})["linkNow"] != nil && resultdat["linkData"].(map[string]interface{})["linkNow"].(bool) == true {
		resultdat["linkData"].(map[string]interface{})["linkNow"] = false
		sync = true
	}

	reference, err := c.executePostJobAndReturnObjectReference("/database/link", resultdat)
	if err != nil {
		return nil, err
	}
	//if "LinkNow" was set to true, then sync (snapshot) the dSource
	if sync == true {
		if err = c.SyncDatabase(reference.(string)); err != nil {
			errorMessage := fmt.Errorf("Unable to sync database: %s", err)
			return nil, errorMessage
		}
	}
	return reference, err
}

// DeleteDSource is a convenience function, as it just invokes DeleteDatabase()
func (c *Client) DeleteDSource(v string) error {
	err := c.DeleteDatabase(strings.ToUpper(v))
	return err
}

// UpdateDSource is a convenience function, as it just invokes UpdateDatabase()
func (c *Client) UpdateDSource(r string, o *OracleDatabaseContainerStruct) error {
	err := c.UpdateDatabase(r, o)
	return err
}

// SyncDatabase performs a snapsync on a database
func (c *Client) SyncDatabase(r string) error {
	url := fmt.Sprintf("/database/%s/sync", strings.ToUpper(r))
	// oracleSyncParameters := OracleSyncParameters{
	// 	Type: "OracleSyncParameters",
	// }
	oracleSyncParameters := CreateOracleSyncParameters(nil, nil, nil, nil, nil)
	err := c.executePostJobAndReturnErrOnly(url, oracleSyncParameters)
	return err
}

//mapifyStruct is a helper function because I couldn't figure out how to reference
//values of nested struct pointers
func mapifyStruct(i interface{}) (map[string]interface{}, error) {
	tbEnc, err := json.Marshal(i)
	fmt.Println(string(tbEnc))
	var resultdat map[string]interface{}
	if err = json.Unmarshal(tbEnc, &resultdat); err != nil { //convert the json to go objects
		return nil, err
	}
	return resultdat, err
}
