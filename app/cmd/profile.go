package cmd

import (
	"fmt"
	"encoding/json"
)

func CmdProfile(url string) error {
	url = "http://" + url + "/profile"
	postData := getPostData()
	fmt.Printf("Post data: %s\n", string(postData))
	responseBody, err := getHttpResponseBody(url, postData)
	fmt.Printf("Response body: %s\n", string(responseBody))
	return err
}

type Message struct {
	Message string `json:"message"`
}

func getPostData() []byte {
	message := Message{
		Message: "test",
	}
	jsonMessage, _ := json.Marshal(message)
	return jsonMessage
}
