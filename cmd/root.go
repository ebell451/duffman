package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

const version = "v1.1.1"

var rootCmd = &cobra.Command{
	Use:   "DuffMan",
	Short: "Fuzzer for Postman collections",
	Long:  `Diagnostic Utility for Fuzzing and Fault Management of API Nodes`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Root().Help()
	},
}

func init() {
	// rootCmd.PersistentFlags().BoolP("verbose", "V", false, "Verbose")
	rootCmd.PersistentFlags().StringP("collection", "f", "", "path to collection file")
	rootCmd.MarkFlagRequired("collection")
	rootCmd.PersistentFlags().StringP("environment", "e", "", "path to environment file")
	rootCmd.Root().CompletionOptions.DisableDefaultCmd = true
}

func Execute() {

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

}
