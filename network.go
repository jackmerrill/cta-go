package ctago

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type NetworkClient struct {
	baseURL url.URL
	apiKey  string
}

func (c *NetworkClient) DoRequest(method string, uri url.URL, body map[string]interface{}) ([]byte, error) {
	httpClient := &http.Client{}
	var jsonBody []byte = nil

	if body != nil {
		j, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		jsonBody = j
	}

	req, err := http.NewRequest(method, uri.String(), bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	q := req.URL.Query()

	q.Add("key", c.apiKey)
	q.Add("outputType", "JSON")

	req.URL.RawQuery = q.Encode()

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bodyBytes, nil
}

func (c *NetworkClient) Parse(data []byte, obj interface{}) error {
	return json.Unmarshal(data, obj)
}

func (c *NetworkClient) Get(uri url.URL, out interface{}) error {
	data, err := c.DoRequest("GET", uri, nil)
	if err != nil {
		return err
	}
	return c.Parse(data, out)
}
