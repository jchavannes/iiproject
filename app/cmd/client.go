package cmd

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"bytes"
	"log"
	"encoding/json"
	"crypto/md5"
	"github.com/jchavannes/iiproject/eid"
)

func loadProfile(url string) {
	url = "http://" + url + "/profile"
	postData := getPostData()
	fmt.Printf("Post data: %s\n", string(postData))
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(postData))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("Response body: %s\n", string(responseBody))
}

type Message struct {
	Message string `json:"message"`
}

func getPostData() []byte {
	message := Message{
		Message: "test",
	}
	jsonMessage, _ := json.Marshal(message)
	fmt.Printf("%x\n", md5.Sum(jsonMessage))
	encrypted, err := eid.EncryptAndSign(jsonMessage)
	if err != nil {
		log.Fatal(err)
	}
	return encrypted
}
