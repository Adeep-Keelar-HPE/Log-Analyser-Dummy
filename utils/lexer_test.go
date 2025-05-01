package utils

import (
	"testing"
)

// Test for Valid input.
func TestLexer_valid(t *testing.T) {
	input := `192.168.1.10 - jane [30/Apr/2025:10:16:00 +0000] "GET /home HTTP/1.1" 302 450 "https://example.com" "Mozilla/5.0"`
	tokens, err := Lexer(input)
	if err != nil {
		t.Errorf("Lexer failed for valid input := %v", err)
	}
	expectedTokenType := []TokenType{
		TokenIP, TokenUser, TokenDate, TokenRequest, TokenStatus, TokenSize, TokenReferrer, TokenUserAgent,
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

// Test for Invalid input.
func TestLexer_invalid(t *testing.T) {
	input := "SOME NONSENSE STRING I HAVE NO IDEA HAHA"
	tokens, err := Lexer(input)
	if err == nil {
		t.Errorf("Lexer should have failed for the invalid input, but returned %v", tokens)
	}
}

// Test for Oversize input.
func TestLexer_oversize(t *testing.T) {
	input := `192.168.1.10 - jane [30/Apr/2025:10:16:00 +0000] "GET /home HTTP/1.1" 302 450 "https://example.com" "Mozilla/5.0" "something extra"`
	tokens, err := Lexer(input)
	if err == nil {
		t.Errorf("Lexer should have failed for the oversize input, but returned %v", tokens)
	} else {
		expectedError := "Invalid number of matches or Empty."
		if err.Error() != expectedError {
			t.Errorf("Expected error message '%s', got '%s'", expectedError, err.Error())
		}
	}
}

// Test for Empty input.
func TestLexer_empty(t *testing.T) {
	input := ""
	tokens, err := Lexer(input)
	if err == nil {
		t.Errorf("Lexer should have failed for the empty input, but returned %v", tokens)
	} else {
		expectedError := "Invalid number of matches or Empty."
		if err.Error() != expectedError {
			t.Errorf("Expected error message '%s', got '%s'", expectedError, err.Error())
		}
	}
}

// Test for One missing field.
func TestLexer_oneMissingField(t *testing.T) {
	input := `192.168.1.10 - jane [30/Apr/2025:10:16:00 +0000] "GET /home HTTP/1.1" 450 "https://example.com" "Mozilla/5.0"`
	tokens, err := Lexer(input)
	if err == nil {
		t.Errorf("Lexer should have failed for the one missing field input, but returned %v", tokens)
	} else {
		expectedError := "Invalid number of matches or Empty."
		if err.Error() != expectedError {
			t.Errorf("Expected error message '%s', got '%s'", expectedError, err.Error())
		}
	}
}
