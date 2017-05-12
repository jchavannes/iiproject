package cmd

import (
	"fmt"
	"github.com/jchavannes/iiproject/eid/client"
	"github.com/jchavannes/iiproject/eid/key"
)

const CliUser = "cli"

func CmdProfile(url string) error {
	profileResponse, err := client.GetProfile(url, "dev2:8252/u/" + CliUser, key.Pair{
		PublicKey: []byte(CliPublicKey),
		PrivateKey: []byte(CliPrivateKey),
	})
	if err != nil {
		return err
	}
	fmt.Printf("profileResponse.Body: %s\n", profileResponse.Body)
	return err
}
