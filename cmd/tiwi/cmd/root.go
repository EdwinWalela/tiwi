package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "tiwi",
	Short: "tiwi generates static sites from markdown",
	Long:  "tiwi parses and builds html sites from markdown",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		log.Fatal("error")
	}
}
