package client

import (
	"github.com/jchavannes/iiproject/eid/api"
	"encoding/json"
	"github.com/jchavannes/iiproject/eid"
)

func GetId(eidUrl string) (*api.IdGetResponse, error) {
	// Execute http request and get response
	url := "https://" + eid.ConvertShortEidUrlIntoFull(eidUrl) + "/id"
	responseBody, err := getResponseJson(url, api.IdRequest{
		Name: "/get",
	})

	// Parse Message
	var idGetResponse api.IdGetResponse
	err = json.Unmarshal(responseBody, &idGetResponse)
	if err != nil {
		return nil, err
	}

	return &idGetResponse, nil
}
