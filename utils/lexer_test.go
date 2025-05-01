package utils

import (
	"testing"
)

func TestLexer_valid(t *testing.T) {
	input := `192.168.1.10 - jane [30/Apr/2025:10:16:00 +0000] "GET /home HTTP/1.1" 302 450 "https://example.com" "Mozilla/5.0"`
	tokens, err := Lexer(input)
	if err != nil {
		t.Errorf("Lexer failed for valid input := %v", err)
	}
	expectedTokenType := []TokenType{
		TokenIP, TokenUser, TokenDate, TokenRequest, TokenStatus, TokenSize, TokenReferrer, TokenUserAgent, TokenErrorCode,
	}

	if len(tokens) != len(expectedTokenType) {
		t.Errorf("Lexer returned %d tokens, expected %d", len(tokens), len(expectedTokenType))
	}

	for i, expected := range expectedTokenType {
		if tokens[i].Type != expected {
			t.Errorf("Expected token %d to be %v, got %v", i, expected, tokens[i].Type)
		}
	}

}
