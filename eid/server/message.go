package server

import (
	"github.com/jchavannes/iiproject/eid/api"
	"encoding/json"
	"github.com/jchavannes/iiproject/eid/client"
	"errors"
	"github.com/jchavannes/iiproject/eid"
	"fmt"
	"github.com/jchavannes/go-pgp/pgp"
	"regexp"
)

func ProcessMessageRequest(body []byte, serverPrivateKey eid.KeyPair) (*api.MessageSend, []byte, error) {
	serverPrivateEntity, err := serverPrivateKey.GetPgpEntity()
	if err != nil {
		return nil, []byte{}, fmt.Errorf("Unable to get server private key: %s", err)
	}

	fmt.Printf("body: %s\n", body)
	decrypted, err := pgp.Decrypt(serverPrivateEntity, body)
	if err != nil {
		return nil, []byte{}, fmt.Errorf("Unable to decrypt message: %s", err)
	}

	reg := regexp.MustCompile(`-----BEGIN PGP SIGNATURE-----[A-Za-z0-9/=+\s]+-----END PGP SIGNATURE-----$`)
	signature := reg.FindString(string(decrypted))
	jsonString := reg.ReplaceAllString(string(decrypted), "")

	var messageSend api.MessageSend
	err = json.Unmarshal([]byte(jsonString), &messageSend)
	if err != nil {
		return nil, []byte{}, fmt.Errorf("Error unmarshalling message send: %s", err)
	}

	idGetResponse, err := client.GetId(messageSend.Eid)
	if err != nil {
		return nil, []byte{}, fmt.Errorf("Unable to get client eid: %s", err)
	}
	clientPublicKey := eid.KeyPair{
		PublicKey: []byte(idGetResponse.PublicKey),
	}

	clientPublicEntity, err := clientPublicKey.GetPgpEntity()
	if err != nil {
		return nil, []byte{}, fmt.Errorf("Error getting client public entity: %s", err)
	}

	err = pgp.Verify(clientPublicEntity, []byte(jsonString), []byte(signature))
	if err != nil {
		return nil, []byte{}, fmt.Errorf("Unable to verify message request: %s", err)
	}

	switch messageSend.Name {
	case "/send":
		messageResponse := api.MessageSendResponse{
			Acknowledged: true,
		}
		encrypted, err := eid.JsonMarshalSignAndEncrypt(messageResponse, clientPublicKey, serverPrivateKey)
		if err != nil {
			return nil, []byte{}, fmt.Errorf("Error encrypting response: %s", err)
		}

		return &messageSend, encrypted, nil
	}

	return nil, []byte{}, errors.New("Unprocessable entity")
}
