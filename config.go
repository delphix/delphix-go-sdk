package delphix

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	resty "gopkg.in/resty.v1"
)

//return an APISession object
//v = APIVersion Struct
//l = Locale as an IETF BCP 47 language tag, defaults to 'en-US'.
//c = Client software identification token.
func CreateAPISession(v APIVersion, l string, c string) (APISession, error) {
	if l == "" {
		l = "en-US"
	}
	if len(c) > 63 {
		err := fmt.Errorf("Client ID specified cannot be longer than 64 characters.\nYou provided %s", c)
		return APISession{}, err
	}
	apiSession := APISession{
		Type:    "APISession",
		Version: &v,
		Locale:  l,
		Client:  c,
	}
	return apiSession, nil
}

//return an APISession object
func CreateAPIVersion(major int, minor int, micro int) (APIVersion, error) {
	t := new(string)
	maj := new(int)
	min := new(int)
	mic := new(int)
	*t = "APIVersion"
	*maj = major
	*min = minor
	*mic = micro

	apiVersion := APIVersion{
		Type:  t,
		Major: maj,
		Minor: min,
		Micro: mic,
	}
	return apiVersion, nil
}

type AuthSuccess struct {
	ID, Message string
}

type RespError struct {
	Type        string `json:"type,omitempty"`
	Status      string `json:"status,omitempty"`
	ErrorStruct `json:"error,omitempty"`
}

type ErrorStruct struct {
	Type          string `json:"type,omitempty"`
	Details       string `json:"details,omitempty"`
	ID            string `json:"id,omitempty"`
	CommandOutput string `json:"commandOutput,omitempty"`
	diagnosis     string `json:"diagnosis,omitempty"`
}
type Client struct {
	url, username, password string
}

func NewClient(username, password, url string) *Client {
	return &Client{
		url:      url,
		username: username,
		password: password,
	}
}

func (c *Client) LoadAndValidate() error {

	versionStruct, err := CreateAPIVersion(1, 9, 0)
	if err != nil {
		return err
	}
	apiStruct, err := CreateAPISession(versionStruct, "", "")
	if err != nil {
		return err
	}
	resty.DefaultClient.
		SetTimeout(time.Duration(30 * time.Second)).
		SetRetryCount(3).
		SetRetryWaitTime(5 * time.Second).
		SetRetryMaxWaitTime(20 * time.Second)

	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(apiStruct).
		Post(c.url + "/session")

	result := resp.Body()
	var resultdat map[string]interface{}
	if err = json.Unmarshal(result, &resultdat); err != nil { //convert the json to go objects
		return err
	}

	if resultdat["status"].(string) == "ERROR" {
		errorMessage := string(result)
		err = fmt.Errorf(errorMessage)
		if err != nil {
			return err
		}
	}

	resp, err = resty.R().
		SetHeader("Content-Type", "application/json").
		SetResult(AuthSuccess{}).
		SetBody(LoginRequest{
			Type:     "LoginRequest",
			Username: c.username,
			Password: c.password,
		}).
		Post(c.url + "/login")
	if err != nil {
		return err
	}
	if http.StatusOK != resp.StatusCode() {
		err = fmt.Errorf("Delphix Username/Password incorrect")
		if err != nil {
			return err
		}
	}
	result = resp.Body()
	if err = json.Unmarshal(result, &resultdat); err != nil { //convert the json to go objects
		return err
	}

	if resultdat["status"].(string) == "ERROR" {
		errorMessage := string(result)
		err = fmt.Errorf(errorMessage)
		if err != nil {
			return err
		}
	}
	return nil
}
