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
	MINUS
	BANG
	ASTERISK
	SLASH

	EQ
	NEQ
	GT
	LT

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
	RETURN
	TRUE
	FALSE
	IF
	ELSE
)

var reservedChar = map[byte]TokenType{
	'=': ASSIGN,
	'+': PLUS,
	'-': MINUS,
	'!': BANG,
	'*': ASTERISK,
	'/': SLASH,
	'>': GT,
	'<': LT,
	'(': LEFT_PARENTHESIS,
	')': RIGHT_PARENTHESIS,
	'{': LEFT_BRACE,
	'}': RIGHT_BRACE,
	',': COMMA,
	';': SEMICOLON,
	0:   EOF,
}

var reserverMultiChar = map[string]TokenType{
	"==": EQ,
	"!=": NEQ,
}

var reserverKeyword = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
}

type TokenType uint8

type Token struct {
	Type    TokenType
	Literal string
}
