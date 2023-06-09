// Code generated by "stringer -type TokenType"; DO NOT EDIT.

package lexer

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ILLEGAL-0]
	_ = x[EOF-1]
	_ = x[IDENTIFIER-2]
	_ = x[INT-3]
	_ = x[ASSIGN-4]
	_ = x[PLUS-5]
	_ = x[MINUS-6]
	_ = x[BANG-7]
	_ = x[ASTERISK-8]
	_ = x[SLASH-9]
	_ = x[EQ-10]
	_ = x[NEQ-11]
	_ = x[GT-12]
	_ = x[LT-13]
	_ = x[COMMA-14]
	_ = x[SEMICOLON-15]
	_ = x[LEFT_PARENTHESIS-16]
	_ = x[RIGHT_PARENTHESIS-17]
	_ = x[LEFT_BRACE-18]
	_ = x[RIGHT_BRACE-19]
	_ = x[FUNCTION-20]
	_ = x[LET-21]
	_ = x[RETURN-22]
	_ = x[TRUE-23]
	_ = x[FALSE-24]
	_ = x[IF-25]
	_ = x[ELSE-26]
}

const _TokenType_name = "ILLEGALEOFIDENTIFIERINTASSIGNPLUSMINUSBANGASTERISKSLASHEQNEQGTLTCOMMASEMICOLONLEFT_PARENTHESISRIGHT_PARENTHESISLEFT_BRACERIGHT_BRACEFUNCTIONLETRETURNTRUEFALSEIFELSE"

var _TokenType_index = [...]uint8{0, 7, 10, 20, 23, 29, 33, 38, 42, 50, 55, 57, 60, 62, 64, 69, 78, 94, 111, 121, 132, 140, 143, 149, 153, 158, 160, 164}

func (i TokenType) String() string {
	if i >= TokenType(len(_TokenType_index)-1) {
		return "TokenType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TokenType_name[_TokenType_index[i]:_TokenType_index[i+1]]
}
