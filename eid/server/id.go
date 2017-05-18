package server

import (
	"github.com/jchavannes/iiproject/eid/api"
	"errors"
	"encoding/json"
)

func ProcessIdRequest(body []byte, publicKey []byte) (*api.IdGetResponse, error) {
	var idRequest api.IdRequest
	err := json.Unmarshal(body, &idRequest)
	if err != nil {
		return nil, err
	}
	switch idRequest.Name {
	case "/get":
		resp := api.IdGetResponse{
			PublicKey: string(publicKey),
		}
		return &resp, nil
	}
	return nil, errors.New("Unable to process")
}
