package server

import (
	"github.com/jchavannes/iiproject/eid/api"
	"encoding/json"
	"github.com/jchavannes/iiproject/eid/client"
	"io"
	"io/ioutil"
	"errors"
	"github.com/jchavannes/iiproject/eid"
	"fmt"
)

func ProcessProfileRequest(body []byte, serverPrivateKey eid.KeyPair, profileReader io.Reader) ([]byte, error) {

	var profileRequest api.ProfileRequest
	err := json.Unmarshal(body, &profileRequest)
	if err != nil {
		return []byte{}, fmt.Errorf("Error unmarshalling profile body: %s", err)
	}

	switch profileRequest.Name {
	case "/get":
		profile, err := ioutil.ReadAll(profileReader)
		if err != nil {
			return []byte{}, fmt.Errorf("Error reading profile: %s", err)
		}

		idGetResponse, err := client.GetId(profileRequest.Eid)
		if err != nil {
			return []byte{}, fmt.Errorf("Error getting client id response: %s", err)
		}

		profileGetResponse := api.ProfileGetResponse{
			Body: string(profile),
		}
		clientPublicKey := eid.KeyPair{
			PublicKey: []byte(idGetResponse.PublicKey),
		}
		encrypted, err := eid.JsonMarshalSignAndEncrypt(profileGetResponse, clientPublicKey, serverPrivateKey)
		if err != nil {
			return []byte{}, fmt.Errorf("Error json marshalling, signing, and encrypting: %s", err)
		}

		return encrypted, nil
	}

	return []byte{}, errors.New("Unprocessable entity")
}
