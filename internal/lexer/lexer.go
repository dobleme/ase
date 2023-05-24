package lexer

type Lexer struct {
	input        string
	position     int
	readPosition int
	char         byte
}

const (
	RESERVED_CHAR uint8 = iota
	LETTER
	NUMBER

	ILLEGAL_CHAR
)

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()

	return l
}

func (l *Lexer) NextToken() *Token {
	var tType TokenType
	var tLiteral string

	l.skipWhitespace()

	switch charType(l.char) {
	case RESERVED_CHAR:
		if l.char == '=' || l.char == '!' {
			if l.peekChar() == '=' {
				char := l.char
				l.readChar()

				tLiteral = string([]byte{char, l.char})
				tType = reserverMultiChar[tLiteral]
				break
			}
		}

		tType, _ = reservedChar[l.char]
		tLiteral = string(l.char)
		if tType == EOF {
			tLiteral = ""
		}

	case LETTER:
		var ok bool
		tLiteral = l.readIdentifier()
		if tType, ok = reserverKeyword[tLiteral]; !ok {
			tType = IDENTIFIER
		}
		return &Token{Type: tType, Literal: tLiteral}

	case NUMBER:
		tType = INT
		tLiteral = l.readNumber()
		return &Token{Type: tType, Literal: tLiteral}

	default:
		tType = ILLEGAL
		tLiteral = string(l.char)
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

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func (l *Lexer) readIdentifier() string {
	pos := l.position
	for isLetter(l.char) {
		l.readChar()
	}

	return l.input[pos:l.position]
}

func (l *Lexer) readNumber() string {
	pos := l.position
	for isNumber(l.char) {
		l.readChar()
	}

	return l.input[pos:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

func charType(ch byte) uint8 {
	if isReservedChar(ch) {
		return RESERVED_CHAR
	}

	if isLetter(ch) {
		return LETTER
	}

	if isNumber(ch) {
		return NUMBER
	}

	return ILLEGAL_CHAR
}

func isReservedChar(ch byte) bool {
	_, ok := reservedChar[ch]
	return ok
}

func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

func isNumber(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
