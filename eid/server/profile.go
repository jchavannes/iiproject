package server

import (
	"github.com/jchavannes/iiproject/eid/api"
	"encoding/json"
	"github.com/jchavannes/go-pgp/pgp"
	"github.com/jchavannes/iiproject/eid/client"
	"io"
	"io/ioutil"
	"errors"
)

func ProcessProfileRequest(body []byte, publicKey []byte, privateKey []byte, profileReader io.Reader) ([]byte, error) {

	var profileRequest api.ProfileRequest
	err := json.Unmarshal(body, &profileRequest)
	if err != nil {
		return []byte{}, err
	}

	switch profileRequest.Name {
	case "/get":
		profile, err := ioutil.ReadAll(profileReader)
		if err != nil {
			return []byte{}, err
		}
		profileGetResponse := api.ProfileGetResponse{
			Body: string(profile),
		}

		jsonResponse, err := json.Marshal(profileGetResponse)
		if err != nil {
			return []byte{}, err
		}

		privEntity, err := pgp.GetEntity(publicKey, privateKey)
		if err != nil {
			return []byte{}, err
		}

		signature, err := pgp.Sign(privEntity, jsonResponse)
		if err != nil {
			return []byte{}, err
		}

		idGetResponse, err := client.GetId(profileRequest.Eid)
		if err != nil {
			return []byte{}, err
		}

		pubEntity, err := pgp.GetEntity([]byte(idGetResponse.PublicKey), nil)
		if err != nil {
			return []byte{}, err
		}

		message := append(jsonResponse, signature...)
		encrypted, err := pgp.Encrypt(pubEntity, message)
		if err != nil {
			return []byte{}, err
		}

		return encrypted, nil
	}
	return []byte{}, errors.New("Unprocessable entity")
}
