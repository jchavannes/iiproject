package client

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

func getResponseJson(url string, postData interface{}) ([]byte, error) {
	// Convert post data to json
	postDataJson, err := json.Marshal(postData)
	if err != nil {
		return []byte{}, err
	}
	return getResponse(url, postDataJson)
}

func getResponse(url string, data []byte) ([]byte, error) {
	// Create and execute http request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	// Get response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return responseBody, nil
}
