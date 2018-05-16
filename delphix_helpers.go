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

func (c *Client) CreateDatabase(o *OracleProvisionParameters) (
	interface{}, error) {
	reference, err := c.executePostJobAndReturnObjectReference("/database/provision", o)
	return reference, err
}

func (c *Client) UpdateDatabase(r string, o *OracleDatabaseContainer) error {
	url := fmt.Sprintf("/database/%s", r)
	err := c.executePostJobAndReturnErrOnly(url, o)
	return err

}

func (c *Client) DeleteDatabase(v string) error {

	o := OracleDeleteParameters{
		Type: "OracleDeleteParameters",
	}
	url := fmt.Sprintf("/database/%s/delete", strings.ToUpper(v))
	err := c.executePostJobAndReturnErrOnly(url, o)
	return err
}

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

func (c *Client) CreateDomain() (interface{}, error) {
	body, err := CreateDomainJSONBody(c)
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

func (c *Client) WaitForEngineReady(p int, t int) error {
	log.Printf("Waiting up to %v seconds for the DDDP to be ready", t)
	timeOut := 0
	for timeOut < t {
		fmt.Println("Waiting for Delphix DDP")
		time.Sleep(time.Duration(p) * time.Second)
		timeOut = timeOut + p
		if err := c.LoadAndValidate(); err == nil {
			break
		}
	}
	return nil
}

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

func CreateDomainJSONBody(c *Client) (string, error) {
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
					"devices": [
						%s
					]
				}
		`, deviceString), nil
}

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

func (c *Client) CreateEnvironment(h *HostEnvironmentCreateParameters) (
	interface{}, error) {
	url := "/environment"
	reference, err := c.executePostJobAndReturnObjectReference(url, h)

	return reference, err
}

func (c *Client) DeleteEnvironment(r string) error {
	url := fmt.Sprintf("/environment/%s/delete", strings.ToUpper(r))
	err := c.executePostJobAndReturnErrOnly(url, "{}")

	return err
}

func (c *Client) UpdateEnvironment(r string, h *UnixHostEnvironment) error {
	url := fmt.Sprintf("/environment/%s", r)

	err := c.executePostJobAndReturnErrOnly(url, h)

	return err
}

func (c *Client) CreateDSource(l *LinkParameters) (
	interface{}, error) {

	sync := false
	resultdat, err := mapifyStruct(l)
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
	//if "LinkNow" was set to true, then sync (snapshot) the dSource
	if sync == true {
		c.SyncDatabase(reference.(string))
	}
	return reference, err
}

//DeleteDSource is a convenience function, as it just invokes DeleteDatabase()
func (c *Client) DeleteDSource(v string) error {
	err := c.DeleteDatabase(strings.ToUpper(v))
	return err
}

//UpdateDSource is a convenience function, as it just invokes UpdateDatabase()
func (c *Client) UpdateDSource(r string, o *OracleDatabaseContainer) error {
	err := c.UpdateDatabase(r, o)
	return err
}

func (c *Client) SyncDatabase(r string) error {
	url := fmt.Sprintf("/database/%s/sync", strings.ToUpper(r))
	oracleSyncParameters := OracleSyncParameters{
		Type: "OracleSyncParameters",
	}
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
