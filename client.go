package pkge

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type HttpClient struct {
	BaseURL string
	APIKey  string
}

type Error struct {
	Code    int         `json:"code"`
	Payload interface{} `json:"payload"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%d: %v\n", e.Code, e.Payload)
}

func NewHttpClient(baseURL, apiKey string) *HttpClient {
	return &HttpClient{
		BaseURL: baseURL,
		APIKey:  apiKey,
	}
}

func (c *HttpClient) Do(method, endpoint string, params url.Values, result interface{}) error {
	end := c.BaseURL + endpoint + "?" + params.Encode()
	req, err := http.NewRequest(method, end, nil)
	if err != nil {
		return errors.New("Request creation failed: " + err.Error())
	}

	req.Header.Set("X-Api-Key", c.APIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return errors.New("request failed: " + err.Error())
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		var apiError Error
		err := json.Unmarshal(responseBody, &apiError)
		if err != nil {
			return fmt.Errorf("failed to unmarshal error response: %w", err)
		}
		return &apiError
	}

	if err != nil {
		return errors.New("failed to read response body: " + err.Error())
	}
	if err := json.Unmarshal(responseBody, &result); err != nil {
		return errors.New("failed to unmarshal response: " + err.Error())
	}
	return nil
}
