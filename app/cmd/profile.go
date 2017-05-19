package cmd

import (
	"fmt"
	"github.com/jchavannes/iiproject/eid/client"
	"github.com/jchavannes/iiproject/eid"
)

const CliUser = "cli"

func CmdProfile(url string) error {
	profileResponse, err := client.GetProfile(url, "dev2:8252/u/" + CliUser, eid.KeyPair{
		PublicKey: []byte(CliPublicKey),
		PrivateKey: []byte(CliPrivateKey),
	})
	if err != nil {
		return err
	}
	fmt.Printf("profileResponse: %#v\n", profileResponse)
	return err
}
