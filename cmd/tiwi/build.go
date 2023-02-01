package cmd

import (
	"github.com/edwinwalela/tiwi/pkg/parse"
	"github.com/spf13/cobra"
)

var parseCmd = &cobra.Command{
	Use:   "build",
	Short: "Generates HTML files",
	Long:  "Parses markdown and produces HTML code",
	Run: func(cmd *cobra.Command, args []string) {
		parse.Build(args)
	},
}

func init() {
	RootCmd.AddCommand(parseCmd)
}
