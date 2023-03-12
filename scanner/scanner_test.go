package scanner

import (
	"github.com/jmptc/glox/token"
	"testing"
)

func TestScanTokens(t *testing.T) {
	s := NewScanner("(){},.-;/*+")

	expectedTokens := []token.Token{
		{TokType: token.LEFT_PAREN, Lexeme: "("},
		{TokType: token.RIGHT_PAREN, Lexeme: ")"},
		{TokType: token.LEFT_BRACE, Lexeme: "{"},
		{TokType: token.RIGHT_BRACE, Lexeme: "}"},
		{TokType: token.COMMA, Lexeme: ","},
		{TokType: token.DOT, Lexeme: "."},
		{TokType: token.MINUS, Lexeme: "-"},
		{TokType: token.SEMICOLON, Lexeme: ";"},
		{TokType: token.SLASH, Lexeme: "/"},
		{TokType: token.STAR, Lexeme: "*"},
		{TokType: token.PLUS, Lexeme: "+"},
		{TokType: token.EOF, Lexeme: "EOF"},
	}

	tokens := s.ScanTokens()

	for i, expectedTok := range expectedTokens {
		if expectedTok.TokType != tokens[i].TokType {
			t.Fatalf("TokType mismatch at %d index. expected: %s got: %s",
				i, expectedTok.TokType, tokens[i].TokType)
		}

		if expectedTok.Lexeme != tokens[i].Lexeme {
			t.Fatalf("Lexeme mismatch at %d index. expected: %s got %s",
				i, expectedTok.Lexeme, tokens[i].Lexeme)
		}
	}

}
