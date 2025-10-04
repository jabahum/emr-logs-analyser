package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "emr-logs-analyser",
	Short: "A CLI tool for analyzing EMR logs",
	Long:  `A CLI tool for analyzing EMR logs with various features and functionalities.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
