package lexer

import (
	"fmt"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	char         byte
}

var tokenType = map[byte]TokenType{
	'=': ASSIGN,
	'+': PLUS,
	'(': LEFT_PARENTHESIS,
	')': RIGHT_PARENTHESIS,
	'{': LEFT_BRACE,
	'}': RIGHT_BRACE,
	',': COMMA,
	';': SEMICOLON,
	0:   EOF,
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()

	return l
}

func (l *Lexer) NextToken() *Token {
	tType, ok := tokenType[l.char]
	if !ok {
		panic(fmt.Sprintf("Token %v incorrect", l.char))
	}

	tLiteral := string(l.char)
	if tType == EOF {
		tLiteral = ""
	}

	l.readChar()

	return &Token{Type: tType, Literal: tLiteral}
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}
