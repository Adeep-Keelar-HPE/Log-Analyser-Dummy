package utils

import "fmt"

// The Parser function takes a string input, and a map of string keys to string values.
func Parser(input string) {
	tokens, err := Lexer(input)
	if err != nil {
		fmt.Println("Error parsing input:", err)
	}
	fmt.Println(tokens)
	fmt.Println("----------------------")

}
