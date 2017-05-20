package cmd

import (
	"github.com/spf13/cobra"
	"errors"
)

var (
	iiCmd = &cobra.Command{
		Use:   "iiproject",
		Short: "Run iiProject application",
	}

	webCmd = &cobra.Command{
		Use:   "web",
		Short: "Main browser application",
		RunE: func(c *cobra.Command, args []string) error {
			return CmdWeb()
		},
	}

	apiCmd = &cobra.Command{
		Use:   "api",
		Short: "eId API",
		RunE: func(c *cobra.Command, args []string) error {
			return CmdApi()
		},
	}

	profileCmd = &cobra.Command{
		Use:   "profile",
		Short: "Implementation of making a profile API call",
		RunE: func(c *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Must specify an eid.")
			}
			return CmdProfile(args[0])
		},
	}

	idCmd = &cobra.Command{
		Use:   "id",
		Short: "Implementation of making an id API call",
		RunE: func(c *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Must specify an eid.")
			}
			return CmdId(args[0])
		},
	}

	generateKeyPairCmd = &cobra.Command{
		Use:   "generate",
		Short: "Generate and output PGP key pair",
		RunE: func(c *cobra.Command, args []string) error {
			return CmdGenerate()
		},
	}
)

func Run() error {
	iiCmd.AddCommand(webCmd)
	iiCmd.AddCommand(apiCmd)
	iiCmd.AddCommand(profileCmd)
	iiCmd.AddCommand(idCmd)
	iiCmd.AddCommand(generateKeyPairCmd)
	return iiCmd.Execute()
}
