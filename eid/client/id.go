package client

import (
	"github.com/jchavannes/iiproject/eid/api"
	"encoding/json"
	"strings"
	"regexp"
)

func GetId(eidUrl string) (*api.IdGetResponse, error) {
	// Execute http request and get response
	url := "https://" + convertEidUrl(eidUrl) + "/id"
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

func convertEidUrl(eidUrl string) string {
	if ! strings.Contains(eidUrl, "@") {
		return eidUrl
	}

	reg := regexp.MustCompile(`^[^@]+@`)
	username := reg.FindString(eidUrl)
	username = username[:len(username) - 1]
	domain := reg.ReplaceAllString(eidUrl, "")
	realEidUrl := domain + "/id/" + username
	return realEidUrl
}
