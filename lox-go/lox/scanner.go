package lox

import (
	"fmt"
	"strconv"
	"unicode"
)

var keywords = map[string]TokenType{}

func init() {
	keywords["and"] = AND
	keywords["class"] = CLASS
	keywords["else"] = ELSE
	keywords["false"] = FALSE
	keywords["for"] = FOR
	keywords["fun"] = FUN
	keywords["if"] = IF
	keywords["nil"] = NIL
	keywords["or"] = OR
	keywords["print"] = PRINT
	keywords["return"] = RETURN
	keywords["super"] = SUPER
	keywords["this"] = THIS
	keywords["true"] = TRUE
	keywords["var"] = VAR
	keywords["while"] = WHILE
}

type Scanner struct {
	lox *Lox

	source string
	tokens []*Token

	start   int
	current int
	line    int
}

func NewScanner(lox *Lox, source string) Scanner {
	return Scanner{
		lox:    lox,
		source: source,
		line:   1,
	}
}

func (s *Scanner) ScanTokens() []*Token {
	for !s.isAtEnd() && !s.lox.hadError {
		s.start = s.current
		s.scanToken()
	}
	s.tokens = append(s.tokens, &Token{TokenType: EOF, Lexeme: "", Literal: nil, Line: s.line})
	return s.tokens
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case '(':
		s.addToken(LEFT_PAREN, nil)
	case ')':
		s.addToken(RIGHT_PAREN, nil)
	case '{':
		s.addToken(LEFT_BRACE, nil)
	case '}':
		s.addToken(RIGHT_BRACE, nil)
	case ',':
		s.addToken(COMMA, nil)
	case '.':
		s.addToken(DOT, nil)
	case '-':
		s.addToken(MINUS, nil)
	case '+':
		s.addToken(PLUS, nil)
	case ';':
		s.addToken(SEMICOLON, nil)
	case '*':
		s.addToken(STAR, nil)
	case '!':
		t := BANG
		if s.match('=') {
			t = BANG_EQUAL
		}
		s.addToken(t, nil)
	case '=':
		t := EQUAL
		if s.match('=') {
			t = EQUAL_EQUAL
		}
		s.addToken(t, nil)
	case '<':
		t := LESS
		if s.match('=') {
			t = LESS_EQUAL
		}
		s.addToken(t, nil)
	case '>':
		t := GREATER
		if s.match('=') {
			t = GREATER_EQUAL
		}
		s.addToken(t, nil)
	case '/':
		if s.match('/') {
			for s.peek() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(SLASH, nil)
		}
	case ' ':
	case '\r':
	case '\t':
	case '\n':
		s.line++
	case '"':
		s.string()
	default:
		if isDigit(c) {
			s.number()
		} else if isAlpha(c) {
			s.identifier()
		} else {
			s.lox.Error(s.line, "Unexpected character.")
		}
	}
}

func (s *Scanner) advance() byte {
	s.current++
	return s.source[s.current-1]
}

func (s *Scanner) addToken(tokenType TokenType, literal any) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, &Token{TokenType: tokenType, Lexeme: text, Literal: literal, Line: s.line})
}

func (s *Scanner) match(c byte) bool {
	if s.isAtEnd() || c != s.source[s.current] {
		return false
	}

	s.current++
	return true
}

func (s *Scanner) peek() byte {
	if s.isAtEnd() {
		return 0
	}
	return s.source[s.current]
}

func (s *Scanner) peekNext() byte {
	if s.current+1 >= len(s.source) {
		return 0
	}
	return s.source[s.current+1]
}

func isDigit(c byte) bool {
	return unicode.IsDigit(rune(c))
}

func isAlpha(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z' || c == '_'
}

func isAlphaNumeric(c byte) bool {
	return isAlpha(c) || isDigit(c)
}

func (s *Scanner) string() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' { //支持多行字符串
			s.line++
		}
		s.advance()
	}

	if s.isAtEnd() {
		s.lox.Error(s.line, "Unterminated string.")
		return
	}
	s.advance()

	val := s.source[s.start+1 : s.current-1]
	s.addToken(STRING, val)
}

func (s *Scanner) number() {
	for isDigit(s.peek()) {
		s.advance()
	}
	if s.peek() == '.' && isDigit(s.peekNext()) {
		s.advance()
		for isDigit(s.peek()) {
			s.advance()
		}
	}

	val, err := strconv.ParseFloat(s.source[s.start:s.current], 64)
	if err != nil {
		s.lox.Error(s.line, fmt.Sprintf("wrong float format: %s", s.source[s.start:s.current]))
		return
	}

	s.addToken(NUMBER, val)
}

func (s *Scanner) identifier() {
	for isAlphaNumeric(s.peek()) {
		s.advance()
	}
	text := s.source[s.start:s.current]
	tokenType := IDENTIFIER
	if t, ok := keywords[text]; ok {
		tokenType = t
	}

	s.addToken(tokenType, nil)
}
