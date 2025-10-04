package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of emr-logs-analyser",
	Long:  `All software has versions. This is emr-logs-analyser's`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Printf("emr-logs-analyser version %s\n", "1.0.0")
	},
}
