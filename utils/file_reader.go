package utils

import (
	"bufio"
	"log"
	"os"
)

// Helper function to read the contents of the file.
func ReadFileLine(file_path string) ([]string, error) {
	var lines []string
	file, err := os.Open(file_path)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	const maxCapacity int = 1000000
	buf := make([]byte, maxCapacity)
	scanner.Buffer(buf, maxCapacity)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
		return nil, err
	}
	return lines, nil
}
