package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represent the base command.
var rootCmd = &cobra.Command{
	Use: "Analyser",
	Short: "Simple tool to analyse the logs",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {

}
