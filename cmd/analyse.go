package cmd

import (
	"LogAnalyser/pkg"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var analyseCmd = &cobra.Command{
	Use:   "analyse",
	Short: "Analyse the log file and generate the json report.",
	Run: func(cmd *cobra.Command, args []string) {
		input_file, _ := cmd.Flags().GetString("input-file")
		if input_file == "" {
			fmt.Println("please provide the input file path...")
		}
		// Call the function to analyse the log file and generate the json report.
		pkg.AnalyseLogFile(input_file)
	},
}

func init() {
	// Add the analyse command to the root command.
	rootCmd.AddCommand(analyseCmd)

	analyseCmd.Flags().StringP("input-file", "i", viper.GetString("INPUT-FILE"), "Path to the log file")
	analyseCmd.MarkFlagRequired("input-file")

}
