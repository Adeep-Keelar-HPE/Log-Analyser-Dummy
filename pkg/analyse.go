package pkg

import (
	utils "LogAnalyser/utils"
	"fmt"
	"os"
)

func AnalyseLogFile(inputFile string) {
	isvalidFile := utils.IsValidFile(inputFile)
	if isvalidFile {
		lines, err := utils.ReadFileLine(inputFile)
		if err != nil {
			fmt.Println("Error reading file:", err)
			os.Exit(1)
		}
		for _, line := range lines {
			// Parse each line using the Parser function.
			utils.Parser(line)
		}
	} else {
		fmt.Println("File is not valid")
		os.Exit(1)
	}
}
