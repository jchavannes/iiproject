package client

import (
	"github.com/jchavannes/iiproject/eid/api"
	"encoding/json"
	"github.com/jchavannes/iiproject/eid/key"
	"github.com/jchavannes/go-pgp/pgp"
)

func GetProfile(eidUrl string, clientEid string, clientKey key.Pair) (*api.ProfileGetResponse, error) {
	// Execute http request and get response
	url := "http://" + eidUrl + "/profile"
	responseBody, err := getResponse(url, api.ProfileRequest{
		Name: "/get",
		Eid: clientEid,
	})
	if err != nil {
		return nil, err
	}

	pgpEntity, err := clientKey.GetPgpEntity()
	if err != nil {
		return nil, err
	}

	decrypted, err := pgp.Decrypt(pgpEntity, responseBody)
	if err != nil {
		return nil, err
	}

	// Parse Message
	var profileGetResponse api.ProfileGetResponse
	err = json.Unmarshal(decrypted, &profileGetResponse)
	if err != nil {
		return nil, err
	}

	return &profileGetResponse, nil
}
