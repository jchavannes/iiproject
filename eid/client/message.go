package client

import (
	"github.com/jchavannes/iiproject/eid/api"
	"github.com/jchavannes/iiproject/eid"
	"time"
	"fmt"
	"errors"
)

func SendMessage(recipientEid string, senderEid string, message string, serverPublicKey eid.KeyPair, clientPrivateKey eid.KeyPair) (*api.MessageSend, error) {
	messageSend := &api.MessageSend{
		Name: "/send",
		Eid: senderEid,
		Message: message,
		SendTime: time.Now(),
	}

	encrypted, err := eid.JsonMarshalSignAndEncrypt(messageSend, serverPublicKey, clientPrivateKey)
	if err != nil {
		return nil, fmt.Errorf("Error json marshalling, signing, and encrypting message: %s", err)
	}

	fmt.Printf("Encrypted: %s\n", encrypted)

	// Execute http request and get response
	url := "https://" + convertEidUrl(recipientEid) + "/message"
	getMessageSendResponse, err := getResponse(url, encrypted)
	if err != nil {
		return nil, fmt.Errorf("Error getting profile response: %s", err)
	}
	fmt.Printf("url: %s\n", url)
	fmt.Printf("getMessageSendResponse: %s\n", getMessageSendResponse)

	var messageSendResponse api.MessageSendResponse
	err = eid.DecryptVerifyAndUnmarshal(getMessageSendResponse, clientPrivateKey, serverPublicKey, &messageSendResponse)
	if err != nil {
		return nil, fmt.Errorf("Error decrypting, verifying, and json unmarshalling message: %s", err)
	}

	if ! messageSendResponse.Acknowledged {
		return nil, errors.New("Message send not ackowledged.")
	}

	return messageSend, nil
}
