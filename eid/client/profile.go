package client

import (
	"github.com/jchavannes/iiproject/eid/api"
	"encoding/json"
	"github.com/jchavannes/go-pgp/pgp"
	"regexp"
	"github.com/jchavannes/iiproject/eid"
	"fmt"
)

func GetProfile(eidUrl string, clientEid string, clientKey eid.KeyPair) (*api.ProfileGetResponse, error) {
	// Execute http request and get response
	url := "https://" + convertEidUrl(eidUrl) + "/profile"
	getProfileResponseBody, err := getResponseJson(url, api.ProfileRequest{
		Name: "/get",
		Eid: clientEid,
	})
	if err != nil {
		return nil, fmt.Errorf("Error getting profile response: %s", err)
	}

	pgpEntity, err := clientKey.GetPgpEntity()
	if err != nil {
		return nil, fmt.Errorf("Error creating client pgp entity: %s", err)
	}

	decrypted, err := pgp.Decrypt(pgpEntity, getProfileResponseBody)
	if err != nil {
		return nil, fmt.Errorf("Error decrypting profile response: %s", err)
	}

	reg := regexp.MustCompile(`-----BEGIN PGP SIGNATURE-----[A-Za-z0-9/=+\s]+-----END PGP SIGNATURE-----$`)
	signature := reg.FindString(string(decrypted))
	jsonString := reg.ReplaceAllString(string(decrypted), "")

	idGetResponse, err := GetId(eidUrl)
	if err != nil {
		return nil, fmt.Errorf("Error getting id response: %s", err)
	}

	publicKeyEntity, err := pgp.GetEntity([]byte(idGetResponse.PublicKey), []byte{})
	if err != nil {
		return nil, fmt.Errorf("Error creating public key entity: %s", err)
	}

	err = pgp.Verify(publicKeyEntity, []byte(jsonString), []byte(signature))
	if err != nil {
		return nil, fmt.Errorf("Error verifying profile message: %s", err)
	}

	var profileGetResponse api.ProfileGetResponse
	err = json.Unmarshal([]byte(jsonString), &profileGetResponse)
	if err != nil {
		return nil, fmt.Errorf("Error unmarshalling profile response data: %s", err)
	}

	return &profileGetResponse, nil
}
