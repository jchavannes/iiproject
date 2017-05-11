package cmd

import (
	"fmt"
	"encoding/json"
	"github.com/jchavannes/iiproject/eid"
)

func CmdId(url string) error {
	url = "http://" + url + "/id"
	postData := getIdPostDate()
	responseBody, err := getHttpResponseBody(url, postData)
	fmt.Printf("Response body: %s\n", string(responseBody))
	return err
}

func getIdPostDate() []byte {
	req := eid.IdRequest{
		Name: "/get",
	}
	reqByte, err := json.Marshal(req)
	if err != nil {
		return []byte{}
	}
	return reqByte
}