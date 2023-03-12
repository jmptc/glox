package scanner

import (
	"github.com/jmptc/glox/token"
)

type Scanner struct {
	source  string
	tokens  []token.Token
	line    int
	start   int
	current int
}

func NewScanner(source string) Scanner {
	return Scanner{
		source:  source,
		line:    0,
		start:   0,
		current: 0,
	}
}

func (s *Scanner) ScanTokens() []token.Token {
	for !s.atEnd() {
		s.scanToken()
	}

	s.createToken("EOF", token.EOF)
	return s.tokens
}

func (s *Scanner) scanToken() {
	c := s.advance()

	switch c {
	case '(':
		s.createToken("(", token.LEFT_PAREN)
	case ')':
		s.createToken(")", token.RIGHT_PAREN)
	case '{':
		s.createToken("{", token.LEFT_BRACE)
	case '}':
		s.createToken("}", token.RIGHT_BRACE)
	case ',':
		s.createToken(",", token.COMMA)
	case '.':
		s.createToken(".", token.DOT)
	case '-':
		s.createToken("-", token.MINUS)
	case ';':
		s.createToken(";", token.SEMICOLON)
	case '/':
		s.createToken("/", token.SLASH)
	case '*':
		s.createToken("*", token.STAR)
	case '+':
		s.createToken("+", token.PLUS)
	}

}

func (s *Scanner) advance() byte {
	c := s.source[s.current]
	s.current += 1
	return c
}

func (s *Scanner) atEnd() bool {
	if len(s.source) == s.current {
		return true
	}
	return false
}

func (s *Scanner) createToken(literal string, tokType token.TokenType) {
	t := token.Token{TokType: tokType, Lexeme: literal}
	s.tokens = append(s.tokens, t)
}
