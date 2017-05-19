package cmd

import (
	"fmt"
	"github.com/jchavannes/iiproject/eid/client"
)

func CmdId(url string) error {
	idGetResponse, err := client.GetId(url)
	if err == nil {
		fmt.Printf("idGetResponse: %#v\n", idGetResponse)
	}
	return err
}
