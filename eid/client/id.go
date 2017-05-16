package client

import (
	"github.com/jchavannes/iiproject/eid/api"
	"encoding/json"
)

func GetId(eidUrl string) (*api.IdGetResponse, error) {
	// Execute http request and get response
	url := "http://" + eidUrl + "/id"
	responseBody, err := getResponse(url, api.IdRequest{
		Name: "/get",
	})

	// Parse Message
	var idGetResponse api.IdGetResponse
	err = json.Unmarshal(responseBody, &idGetResponse)
	if err!= nil {
		return nil, err
	}

	return &idGetResponse, nil
}
