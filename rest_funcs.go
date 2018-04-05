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
	if err = json.Unmarshal(responseBody, &resultMap); err != nil { //convert the json to go objects
		return nil, err
	}

	return resultMap, err //return the query results
}
