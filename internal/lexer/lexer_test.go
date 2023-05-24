package lexer_test

import (
	"testing"

	"github.com/dobleme/ase-interpreter/internal/lexer"
)

type TokenTest struct {
	expectedType    lexer.TokenType
	expectedLiteral string
}

type TestCase struct {
	code   string
	result []TokenTest
}

var tests = map[string]TestCase{
	"simple shit": {
		code: `=+(){},;`,
		result: []TokenTest{
			{lexer.ASSIGN, "="},
			{lexer.PLUS, "+"},
			{lexer.LEFT_PARENTHESIS, "("},
			{lexer.RIGHT_PARENTHESIS, ")"},
			{lexer.LEFT_BRACE, "{"},
			{lexer.RIGHT_BRACE, "}"},
			{lexer.COMMA, ","},
			{lexer.SEMICOLON, ";"},
			{lexer.EOF, ""},
		},
	},
	"more complex shit": {
		code: `let five = 5;
let ten = 10;
let add = fn(a, b) {
    a + b;
};
let result = add(five, ten);
`,
		result: []TokenTest{
			{lexer.LET, "let"},
			{lexer.IDENTIFIER, "five"},
			{lexer.ASSIGN, "="},
			{lexer.INT, "5"},
			{lexer.SEMICOLON, ";"},
			{lexer.LET, "let"},
			{lexer.IDENTIFIER, "ten"},
			{lexer.ASSIGN, "="},
			{lexer.INT, "10"},
			{lexer.SEMICOLON, ";"},
			{lexer.LET, "let"},
			{lexer.IDENTIFIER, "add"},
			{lexer.ASSIGN, "="},
			{lexer.FUNCTION, "fn"},
			{lexer.LEFT_PARENTHESIS, "("},
			{lexer.IDENTIFIER, "a"},
			{lexer.COMMA, ","},
			{lexer.IDENTIFIER, "b"},
			{lexer.RIGHT_PARENTHESIS, ")"},
			{lexer.LEFT_BRACE, "{"},
			{lexer.IDENTIFIER, "a"},
			{lexer.PLUS, "+"},
			{lexer.IDENTIFIER, "b"},
			{lexer.SEMICOLON, ";"},
			{lexer.RIGHT_BRACE, "}"},
			{lexer.SEMICOLON, ";"},
			{lexer.LET, "let"},
			{lexer.IDENTIFIER, "result"},
			{lexer.ASSIGN, "="},
			{lexer.IDENTIFIER, "add"},
			{lexer.LEFT_PARENTHESIS, "("},
			{lexer.IDENTIFIER, "five"},
			{lexer.COMMA, ","},
			{lexer.IDENTIFIER, "ten"},
			{lexer.RIGHT_PARENTHESIS, ")"},
			{lexer.SEMICOLON, ";"},
		},
	},
	"more simple shit": {
		code: `!-/*5;
5 < 10 > 5;
`,
		result: []TokenTest{
			{lexer.BANG, "!"},
			{lexer.MINUS, "-"},
			{lexer.SLASH, "/"},
			{lexer.ASTERISK, "*"},
			{lexer.INT, "5"},
			{lexer.SEMICOLON, ";"},
			{lexer.INT, "5"},
			{lexer.LT, "<"},
			{lexer.INT, "10"},
			{lexer.GT, ">"},
			{lexer.INT, "5"},
			{lexer.SEMICOLON, ";"},
		},
	},
}

func TestLexer(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			lex := lexer.New(test.code)

			for i, tc := range test.result {
				token := lex.NextToken()

				if token.Type != tc.expectedType {
					t.Fatalf(
						"[pos %d] expect token type %s got %s (%v)",
						i, tc.expectedType, token.Type, token.Literal,
					)
				}

				if token.Literal != tc.expectedLiteral {
					t.Fatalf(
						"[pos %d] expect token literal %q got %q",
						i, tc.expectedLiteral, token.Literal,
					)
				}
			}
		})
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
