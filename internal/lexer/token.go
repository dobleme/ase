package lexer

const (
	ILLEGAL TokenType = iota
	EOF

	// Identifiers + literal
	IDENTIFIER
	INT

	// Operators
	ASSIGN
	PLUS

	// Delimiters
	COMMA
	SEMICOLON
	LEFT_PARENTHESIS
	RIGHT_PARENTHESIS
	LEFT_BRACE
	RIGHT_BRACE

	// Keywords
	FUNCTION
	LET
)

type TokenType uint8

type Token struct {
	Type    TokenType
	Literal string
}
