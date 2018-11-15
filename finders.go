package delphix

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	resty "gopkg.in/resty.v1"
)

// FindSourceByName returns the reference of the named source(n)
func (c *Client) FindSourceByName(n string) (interface{}, error) {
	var err error
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		Get(c.url + "/database") //grab all the databases
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

	databases := resp.Body()

	var dat map[string]interface{}
	if err = json.Unmarshal(databases, &dat); err != nil { //convert the json to go objects
		return nil, err
	}
	results := dat["result"].([]interface{}) //grab the query results
	for _, result := range results {         //loop through the repos
		name := result.(map[string]interface{})["name"]           //grab the repo name
		reference := result.(map[string]interface{})["reference"] //grab the repo UUID
		if name == n {                                            //if the name matches our specified repo
			return reference, nil //return the reference
		}
	}
	log.Printf("Was unable to find a database container with the name of %s in %s", n, string(resp.Body()))
	return nil, nil
}

// FindDatabaseByReference returns the database object (interface) of the reference (r) database
func (c *Client) FindDatabaseByReference(r string) (interface{}, error) {
	var err error
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		Get(c.url + "/database") //grab all the databases
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

	databases := resp.Body()

	var dat map[string]interface{}
	if err = json.Unmarshal(databases, &dat); err != nil { //convert the json to go objects
		return nil, err
	}
	results := dat["result"].([]interface{}) //grab the query results
	for _, result := range results {         //loop through the databases
		reference := result.(map[string]interface{})["reference"] //grab the db UUID
		if reference == r {                                       //if the name matches our specified vdb
			return result.(map[string]interface{}), nil //return the reference
		}
	}
	log.Printf("Was unable to find vdb with a reference of %s in %s", r, string(resp.Body()))
	return nil, nil
}

// FindDatabaseByName returns the database object (interface) of the named (n) database
func (c *Client) FindDatabaseByName(n string) (interface{}, error) {
	var err error
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		Get(c.url + "/database") //grab all the databases
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

	databases := resp.Body()

	var dat map[string]interface{}
	if err = json.Unmarshal(databases, &dat); err != nil { //convert the json to go objects
		return nil, err
	}
	results := dat["result"].([]interface{}) //grab the query results
	for _, result := range results {         //loop through the databases
		name := result.(map[string]interface{})["name"]           //grab the db name
		reference := result.(map[string]interface{})["reference"] //grab the db UUID
		if name == n {                                            //if the name matches our specified vdb
			return reference, nil //return the reference
		}
	}
	log.Printf("Was unable to find vdb with a name of %s in %s", n, string(resp.Body()))
	return nil, nil
}

// FindGroupByName returns the group object (interface) of the named (n) group
func (c *Client) FindGroupByName(n string) (interface{}, error) {
	var err error
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		Get(c.url + "/group") //grab all the groups
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

	groups := resp.Body()

	var dat map[string]interface{}
	if err = json.Unmarshal(groups, &dat); err != nil { //convert the json to go objects
		return nil, err
	}
	results := dat["result"].([]interface{}) //grab the query results
	for _, result := range results {         //loop through the groups
		name := result.(map[string]interface{})["name"]           //grab the group name
		reference := result.(map[string]interface{})["reference"] //grab the group UUID
		if name == n {                                            //if the name matches our specified group
			return reference, nil //return the reference
		}
	}
	log.Printf("Was unable to find group with a name of %s in %s", n, string(resp.Body()))
	return nil, nil
}

// FindRepoReferenceByEnvironmentRefAndOracleHome returns the reference of the Repository by
// environment reference (e) and oracle home (oh) path
func (c *Client) FindRepoReferenceByEnvironmentRefAndOracleHome(e string, oh string) (interface{}, error) {
	var err error
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		Get(c.url + "/repository?environment=" + e) //grab all the repos for the environment
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

	repos := resp.Body()

	var dat map[string]interface{}
	if err = json.Unmarshal(repos, &dat); err != nil { //convert the json to go objects
		return nil, err
	}
	results := dat["result"].([]interface{}) //grab the query results
	for _, result := range results {         //loop through the repos
		name := result.(map[string]interface{})["name"]           //grab the repo name
		reference := result.(map[string]interface{})["reference"] //grab the repo UUID
		if name == oh {                                           //if the name matches our specified repo
			return reference, nil //return the reference
		}
	}
	log.Printf("Was unable to find %s on %s in %s", oh, e, string(resp.Body()))
	return nil, nil
}

// FindEnvironmentByName returns the environment object (interface) of the named (n) environment
func (c *Client) FindEnvironmentByName(n string) (interface{}, error) {
	var err error
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		Get(c.url + "/environment") //grab all the environments
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

	groups := resp.Body()

	var dat map[string]interface{}
	if err = json.Unmarshal(groups, &dat); err != nil { //convert the json to go objects
		return nil, err
	}
	results := dat["result"].([]interface{}) //grab the query results
	for _, result := range results {         //loop through the environments
		name := result.(map[string]interface{})["name"] //grab the environment name
		if n == name {                                  //if the name matches our specified environment
			return result, nil //return the environment object
		}
	}
	log.Printf("Was unable to find environment \"%s\" in %s", n, string(resp.Body()))
	return nil, nil
}

// FindEnvironmentByReference returns the environment object (interface) of the referenced (r) environment
func (c *Client) FindEnvironmentByReference(r string) (interface{}, error) {
	var err error
	url := fmt.Sprintf("/environment/%s", r)
	envObj, err := c.executeReadAndReturnObject(url)
	if err != nil {
		return nil, err
	}
	return envObj, nil
}

// FindUserByName returns the user object (interface) of the named (n) user
func (c *Client) FindUserByName(n string) (interface{}, error) {
	var err error
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		Get(c.url + "/user") //grab all the users
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

	users := resp.Body()

	var dat map[string]interface{}
	if err = json.Unmarshal(users, &dat); err != nil { //convert the json to go objects
		return nil, err
	}
	results := dat["result"].([]interface{}) //grab the query results
	for _, result := range results {         //loop through the users
		name := result.(map[string]interface{})["name"]           //grab the users name
		reference := result.(map[string]interface{})["reference"] //grab the user UUID
		if name == n {                                            //if the name matches our specified user
			return reference, nil //return the reference
		}
	}
	fmt.Printf("Was unable to find %s in %s", n, string(resp.Body()))
	return nil, nil
}

// FindSourceConfigByNameAndRepoReference returns the sourceconfig object (interface) of the
// named (n) sourceconfig for the specified repo reference (r)
func (c *Client) FindSourceConfigByNameAndRepoReference(n string, r string) (interface{}, error) {
	url := fmt.Sprintf("/sourceconfig?repository=%s", r)
	results, err := c.executeListAndReturnResults(url) //grab the query results
	if err != nil {
		return nil, err
	}
	reference, err := returnObjectReferenceByName(results, n)
	if err != nil {
		return nil, err
	}
	return reference, err
}

func returnObjectReferenceByName(o []interface{}, n string) (string, error) {
	obj, err := returnObjectByName(o, n)
	if err != nil {
		return "", err
	}
	reference, ok := obj.(map[string]interface{})["reference"].(string) //grab the object reference
	if !ok {
		return "", fmt.Errorf("Did not find object named %s in %v", n, o)
	}
	return reference, err
}

func returnObjectByName(o []interface{}, n string) (interface{}, error) {
	for _, result := range o { //loop through the objects
		name := result.(map[string]interface{})["name"].(string) //grab the object name
		if name == n {                                           //if the name matches
			return result, nil //return the object
		}
	}
	return nil, fmt.Errorf("Could not find object named %s in %v", n, o)
}

func returnObjectByReference(o []interface{}, r string) (interface{}, error) {
	for _, result := range o { //loop through the objects
		reference := result.(map[string]interface{})["reference"].(string) //grab the object reference
		if reference == r {                                                //if the reference matches
			return result, nil //return the object
		}
	}
	return nil, fmt.Errorf("Could not find object reference %s in %v", r, o)
}

// FindEnvironmentUserByNameAndEnvironmentReference returns the reference of the environment user by
// environment user name (n) and environment reference (r)
func (c *Client) FindEnvironmentUserByNameAndEnvironmentReference(n string, r string) (interface{}, error) {
	url := fmt.Sprintf("/environment/user?environment=%s", r)
	results, err := c.executeListAndReturnResults(url)
	if err != nil {
		return nil, err
	}
	reference, err := returnObjectReferenceByName(results, n)
	if err != nil {
		return nil, err
	}
	return reference, err
}

// FindObjectByReference returns the referenced (r) object type (u)
func (c *Client) FindObjectByReference(u string, r string) (interface{}, error) {
	url := fmt.Sprintf("%s/%s", u, r)
	obj, err := c.executeReadAndReturnObject(url)
	if err != nil {
		return nil, err
	}
	return obj, err
}

// FindObjectByName returns the named (n) object type (u)
func (c *Client) FindObjectByName(u string, n string) (interface{}, error) {
	url := u
	results, err := c.executeListAndReturnResults(url)
	if err != nil {
		return nil, err
	}
	obj, err := returnObjectByName(results, n)
	if err != nil {
		return nil, err
	}
	return obj, err
}

// ReturnSshPublicKey returns the SSH Public Key of the Delphix Engine
func (c *Client) ReturnSshPublicKey() (string, error) {
	systemObj, err := c.executeReadAndReturnObject("/system")
	if err != nil {
		return "", err
	}
	k, ok := systemObj.(map[string]interface{})["sshPublicKey"].(string) //grab the object reference
	if !ok {
		return "", fmt.Errorf("Did not find the sshPublicKey. Something went terribly wrong")
	}
	return k, err

}
