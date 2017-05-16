package client

import (
	"github.com/jchavannes/iiproject/eid/api"
	"encoding/json"
	"github.com/jchavannes/iiproject/eid/key"
	"github.com/jchavannes/go-pgp/pgp"
	"regexp"
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

	reg := regexp.MustCompile(`-----BEGIN PGP SIGNATURE-----[A-Za-z0-9/=+\s]+-----END PGP SIGNATURE-----$`)
	signature := reg.FindString(string(decrypted))
	jsonString := reg.ReplaceAllString(string(decrypted), "")

	idGetResponse, err := GetId(eidUrl)
	if err != nil {
		return nil, err
	}

	publicKeyPacket, err := pgp.GetPublicKeyPacket([]byte(idGetResponse.PublicKey))
	if err != nil {
		return nil, err
	}

	err = pgp.Verify(publicKeyPacket, []byte(jsonString), []byte(signature))
	if err != nil {
		return nil, err
	}

	var profileGetResponse api.ProfileGetResponse
	err = json.Unmarshal([]byte(jsonString), &profileGetResponse)
	if err != nil {
		return nil, err
	}

	return &profileGetResponse, nil
}
