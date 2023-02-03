package cmd

import (
	"github.com/edwinwalela/tiwi/pkg/watch"
	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Build and live reload HTML",
	Long:  "Watches markdown files and rebuilds HTML on file change",
	Run: func(cmd *cobra.Command, args []string) {
		watch.Watch(args)
	},
}

func init() {
	RootCmd.AddCommand(watchCmd)
}
