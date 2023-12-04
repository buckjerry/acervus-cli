package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "acervus",
	Short: "CLI application to interact with Acervus Cloud",
	Long:  "You can use this application to manage your Acervus account, generate, test and deploy Acervus projects.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
