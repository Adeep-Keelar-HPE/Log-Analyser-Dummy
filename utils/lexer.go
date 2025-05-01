package utils

import (
	"fmt"
	"regexp"
)

// Define the token type.
type TokenType string

const (
	TokenIP        TokenType = "IP"
	TokenUser      TokenType = "User"
	TokenDate      TokenType = "Date"
	TokenRequest   TokenType = "Request"
	TokenStatus    TokenType = "Status"
	TokenSize      TokenType = "Size"
	TokenReferrer  TokenType = "Referrer"
	TokenUserAgent TokenType = "UserAgent"
	TokenErrorCode TokenType = "ErrorCode"
)

// Define the Token Struct.
type Token struct {
	Type  TokenType
	Value string
}

func Lexer(input string) ([]Token, error) {
	// The lexer function takes the input string from the parser, and returns a slice of tokens.
	// Each of the tokens will represent a part that will be used in the Analyser.

	// Use a lean regex pattern to populate the tokens.
	pattern := `^(\S+) \S+ (\S+) \[([^\]]+)\] "([^"]+)" (\d{3}) (\d+) "([^"]+)" "([^"]+)"(?: errorCode=(\S+))?$`
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(input)
	if len(matches) == 0 {
		fmt.Println("No matches found")
		return nil, fmt.Errorf("No matches found")
	}
	// Create slice of tokens.
	tokens := []Token{
		{Type: TokenIP, Value: matches[1]},
		{Type: TokenUser, Value: matches[2]},
		{Type: TokenDate, Value: matches[3]},
		{Type: TokenRequest, Value: matches[4]},
		{Type: TokenStatus, Value: matches[5]},
		{Type: TokenSize, Value: matches[6]},
		{Type: TokenReferrer, Value: matches[7]},
		{Type: TokenUserAgent, Value: matches[8]},
	}
	// Optional error code
	if len(matches) > 9 && matches[9] != "" {
		tokens = append(tokens, Token{Type: TokenErrorCode, Value: matches[9]})
	}
	return tokens, nil
}
