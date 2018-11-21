package delphix

import (
	"encoding/json"
	"fmt"
	"net/http"

	resty "gopkg.in/resty.v1"
)

func (c *Client) executePostJobAndReturnObjectReference(u string, p interface{}) (
	interface{}, error) {

	postBody := p
	//DEBUG
	tbEnc, err := json.Marshal(postBody)
	fmt.Println(string(tbEnc))
	//DEBUG
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(postBody).
		Post(c.url + u)

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

	reference := resultdat["result"].(string)           //grab the vdb reference
	if jobNumber, ok := resultdat["job"].(string); ok { //grab the job reference
		c.WaitforDelphixJob(jobNumber)
	}

	return reference, err
}

func (c *Client) executePostJobAndReturnErrOnly(u string, p interface{}) error {

	postBody := p
	//DEBUG
	tbEnc, err := json.Marshal(postBody)
	fmt.Println(string(tbEnc))
	//DEBUG
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(postBody).
		Post(c.url + u)

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

	if jobNumber, ok := resultdat["job"].(string); ok { //grab the job reference
		c.WaitforDelphixJob(jobNumber)
	}

	return err
}

func (c *Client) executeListAndReturnResults(u string) (
	[]interface{}, error) {

	obj, err := c.executeGetReturnBody(u)
	if err != nil {
		return nil, err
	}
	return obj["result"].([]interface{}), err //return the query results
}

func (c *Client) executeReadAndReturnObject(u string) (
	interface{}, error) {
	obj, err := c.executeGetReturnBody(u)
	if err != nil {
		//return err, if anything but object does not exist
		if err.Error() != "exception.executor.object.missing" {
			return nil, err
		}
		//object is missing, return nil, not err
		return nil, nil
	}
	return obj["result"].(interface{}), err //return the query result
}

func (c *Client) executeGetReturnBody(u string) (
	map[string]interface{}, error) {
	var resultMap map[string]interface{}
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetError(&RespError{}).
		Get(c.url + u)
	if err != nil {
		return nil, err
	}

	if http.StatusOK != resp.StatusCode() { //check to make sure our query was good
		errorMessage := string(resp.Body())
		err = fmt.Errorf(errorMessage)
		if err != nil {
			if errorID := resp.Error().(*RespError).ErrorStruct.ID; errorID == "exception.executor.object.missing" {
				//object is missing, return that as the error
				return resultMap, fmt.Errorf(errorID)
			}
			return nil, err
		}
	}
	responseBody := resp.Body()
	fmt.Printf("\n\nRESPONSE: %v\n\n", string(responseBody))
	if err = json.Unmarshal(responseBody, &resultMap); err != nil { //convert the json to go objects
		return nil, err
	}

	if resultMap["status"].(string) == "ERROR" {
		errorMessage := string(resp.Body())
		err = fmt.Errorf(errorMessage)
		if err != nil {
			return nil, err
		}
	}
	return resultMap, err //return the query results
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

func (c *Client) executeGetReturnBodyAndCast(u string) (
	interface{}, error) {
	var resultMap map[string]interface{}
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetError(&RespError{}).
		Get(c.url + u)
	if err != nil {
		return nil, err
	}

	if http.StatusOK != resp.StatusCode() { //check to make sure our query was good
		errorMessage := string(resp.Body())
		err = fmt.Errorf(errorMessage)
		if err != nil {
			if errorID := resp.Error().(*RespError).ErrorStruct.ID; errorID == "exception.executor.object.missing" {
				//object is missing, return that as the error
				return resultMap, fmt.Errorf(errorID)
			}
			return nil, err
		}
	}
	responseBody := resp.Body()
	if err = json.Unmarshal(responseBody, &resultMap); err != nil { //convert the json to go objects
		return nil, err
	}

	switch resultMap["type"].(string) {
	case "ErrorResult":
		errorMessage := string(resp.Body())
		err = fmt.Errorf(errorMessage)
		if err != nil {
			return nil, err
		}
	case "OKResult":
		var okResult *OKResultStruct
		if err = json.Unmarshal(responseBody, &okResult); err != nil { //convert the json to go objects
			return nil, err
		}
		return okResult, err
	case "ListResult":
		var listResult *ListResultStruct
		if err = json.Unmarshal(responseBody, &listResult); err != nil { //convert the json to go objects
			return nil, err
		}
		return listResult, err
	case "DataResult":
		var dataResult *DataResultStruct
		if err = json.Unmarshal(responseBody, &dataResult); err != nil { //convert the json to go objects
			return nil, err
		}
		return dataResult, err
	default:
		err = fmt.Errorf("None of the expected Result types were returned. Received: \n%v", resultMap["type"])
	}
	return nil, err
}

func (c *Client) executeListAndCastAndReturnResults(u string) (
	*ListResultStruct, error) {

	obj, err := c.executeGetReturnBodyAndCast(u)
	if err != nil {
		return nil, err
	}
	switch obj.(type) {
	case *ListResultStruct:
		return obj.(*ListResultStruct), err //return the query results
	default:
		return nil, err
	}
}
