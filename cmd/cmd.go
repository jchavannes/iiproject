package cmd

import (
	"github.com/spf13/cobra"
)

var (
	iiCmd = &cobra.Command{
		Use:   "ii",
		Short: "Run ii",
	}

	webCmd = &cobra.Command{
		Use:   "web",
		RunE: func (c *cobra.Command, args []string) error {
			return Web()
		},
	}
)

func Run() error {
	iiCmd.AddCommand(webCmd)
	return iiCmd.Execute()
}
