package cmd

import (
	"github.com/spf13/cobra"
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
)

func Run() error {
	iiCmd.AddCommand(webCmd)
	return iiCmd.Execute()
}
