package cmd

import (
	"github.com/edwinwalela/tiwi/pkg/create"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Initializes the tiwi project",
	Long:  "Creates a directory containing markdown files for your tiwi project",
	Run: func(cmd *cobra.Command, args []string) {
		create.CreateSite()
	},
}

func init() {
	RootCmd.AddCommand(createCmd)
}
