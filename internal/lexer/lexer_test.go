package lexer_test

import (
	"testing"

	"github.com/dobleme/ase-transpiler/internal/lexer"
)

var tests = []struct {
	expectedType    lexer.TokenType
	expectedLiteral string
}{
	{lexer.ASSIGN, "="},
	{lexer.PLUS, "+"},
	{lexer.LEFT_PARENTHESIS, "("},
	{lexer.RIGHT_PARENTHESIS, ")"},
	{lexer.LEFT_BRACE, "{"},
	{lexer.RIGHT_BRACE, "}"},
	{lexer.COMMA, ","},
	{lexer.SEMICOLON, ";"},
	{lexer.EOF, ""},
}

func TestLexer(t *testing.T) {
	lex := lexer.New(`=+(){},;`)
	for i, tc := range tests {
		token := lex.NextToken()

		if token.Type != tc.expectedType {
			t.Fatalf("[pos %d] expect token type %q got %q", i, tc.expectedType, token.Type)
		}

		if token.Literal != tc.expectedLiteral {
			t.Fatalf("[pos %d] expect token literal %q got %q", i, tc.expectedLiteral, token.Literal)
		}
	}
}

func BenchmarkLexer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		lex := lexer.New(`=+(){},;`)
		for range tests {
			lex.NextToken()
		}
	}
}
