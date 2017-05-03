package cmd

import (
	"github.com/spf13/cobra"
	"errors"
	"github.com/jchavannes/iiproject/eid"
)

var (
	iiCmd = &cobra.Command{
		Use:   "iiproject",
		Short: "Run iiProject application",
	}

	webCmd = &cobra.Command{
		Use:   "web",
		RunE: func (c *cobra.Command, args []string) error {
			return runWeb()
		},
	}

	profileCmd = &cobra.Command{
		Use:   "profile",
		RunE: func (c *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Must specify an eid.")
			}
			loadProfile(args[0])
			return nil
		},
	}

	generateKeyPairCmd = &cobra.Command{
		Use:   "generate",
		RunE: func (c *cobra.Command, args []string) error {
			keyPair, err := eid.GenerateKeyPair("Test Key", "test", "test@jasonc.me")
			if err != nil {
				return err
			}
			println(keyPair.PrivateKey)
			println(keyPair.PublicKey)
			return nil
		},
	}
)

func Run() error {
	iiCmd.AddCommand(webCmd)
	iiCmd.AddCommand(profileCmd)
	iiCmd.AddCommand(generateKeyPairCmd)
	return iiCmd.Execute()
}
