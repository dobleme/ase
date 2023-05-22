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
	_ = x[COMMA-6]
	_ = x[SEMICOLON-7]
	_ = x[LEFT_PARENTHESIS-8]
	_ = x[RIGHT_PARENTHESIS-9]
	_ = x[LEFT_BRACE-10]
	_ = x[RIGHT_BRACE-11]
	_ = x[FUNCTION-12]
	_ = x[LET-13]
}

const _TokenType_name = "ILLEGALEOFIDENTIFIERINTASSIGNPLUSCOMMASEMICOLONLEFT_PARENTHESISRIGHT_PARENTHESISLEFT_BRACERIGHT_BRACEFUNCTIONLET"

var _TokenType_index = [...]uint8{0, 7, 10, 20, 23, 29, 33, 38, 47, 63, 80, 90, 101, 109, 112}

func (i TokenType) String() string {
	if i >= TokenType(len(_TokenType_index)-1) {
		return "TokenType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TokenType_name[_TokenType_index[i]:_TokenType_index[i+1]]
}
