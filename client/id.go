package client

import (
	"github.com/jchavannes/iiproject/eid"
	"encoding/json"
)

func GetId(eidUrl string) (*eid.IdGetResponse, error) {
	// Execute http request and get response
	url := "http://" + eidUrl + "/id"
	responseBody, err := getResponse(url, eid.IdRequest{
		Name: "/get",
	})

	// Parse Message
	var idGetResponse *eid.IdGetResponse
	err = json.Unmarshal(responseBody, idGetResponse)
	if err!= nil {
		return nil, err
	}

	return idGetResponse, nil
}
