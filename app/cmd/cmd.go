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
		RunE: func (c *cobra.Command, args []string) error {
			return CmdWeb()
		},
	}

	profileCmd = &cobra.Command{
		Use:   "profile",
		RunE: func (c *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Must specify an eid.")
			}
			return CmdProfile(args[0])
		},
	}

	idCmd = &cobra.Command{
		Use:   "id",
		RunE: func (c *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("Must specify an eid.")
			}
			return CmdId(args[0])
		},
	}

	generateKeyPairCmd = &cobra.Command{
		Use:   "generate",
		RunE: func (c *cobra.Command, args []string) error {
			return CmdGenerate()
		},
	}
)

func Run() error {
	iiCmd.AddCommand(webCmd)
	iiCmd.AddCommand(profileCmd)
	iiCmd.AddCommand(idCmd)
	iiCmd.AddCommand(generateKeyPairCmd)
	return iiCmd.Execute()
}
