package delphix

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	resty "gopkg.in/resty.v1"
)

// CreateAPISession returns an APISession object
//v = APIVersion Struct
//l = Locale as an IETF BCP 47 language tag, defaults to 'en-US'.
//c = Client software identification token.
func CreateAPISession(v APIVersionStruct, l string, c string) (APISessionStruct, error) {
	if l == "" {
		l = "en-US"
	}
	if len(c) > 63 {
		err := fmt.Errorf("Client ID specified cannot be longer than 64 characters.\nYou provided %s", c)
		return APISessionStruct{}, err
	}
	apiSession := APISessionStruct{
		Type:    "APISession",
		Version: &v,
		Locale:  l,
		Client:  c,
	}
	return apiSession, nil
}

//CreateAPIVersion returns an APISession object
func CreateAPIVersion(major int, minor int, micro int) (APIVersionStruct, error) {
	maj := new(int)
	min := new(int)
	mic := new(int)
	t := "APIVersion"
	*maj = major
	*min = minor
	*mic = micro

	apiVersion := APIVersionStruct{
		Type:  t,
		Major: maj,
		Minor: min,
		Micro: mic,
	}
	return apiVersion, nil
}

// AuthSuccess holds the resty response success message
type AuthSuccess struct {
	ID, Message string
}

// RespError holds the resty response failure message
type RespError struct {
	Type        string `json:"type,omitempty"`
	Status      string `json:"status,omitempty"`
	ErrorStruct `json:"error,omitempty"`
}

// ErrorStruct is the struct of a resty error
type ErrorStruct struct {
	Type          string `json:"type,omitempty"`
	Details       string `json:"details,omitempty"`
	ID            string `json:"id,omitempty"`
	CommandOutput string `json:"commandOutput,omitempty"`
	Diagnosis     string `json:"diagnosis,omitempty"`
}

// Client the structure of a client request
type Client struct {
	url, username, password string
}

// NewClient creates a new client object
func NewClient(username, password, url string) *Client {
	return &Client{
		url:      url,
		username: username,
		password: password,
	}
}

// LoadAndValidate establishes a new client connection
func (c *Client) LoadAndValidate() error {

	versionStruct, err := CreateAPIVersion(1, 11, 9)
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
		SetBody(LoginRequestStruct{
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
