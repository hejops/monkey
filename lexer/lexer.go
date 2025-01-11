package lexer

import (
	"unicode"

	"monkey/token"
)

// source code -[lexer]> tokens -[parser]> AST
// whitespace is ignored

// keywords: let, fn
// identifiers: five, ten
// data: 5, 10
// symbols: = ; ( ) , { } +

type Lexer struct {
	input    string
	position int
	c        byte
	nextPos  int // for peeking
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // init
	return l
}

func (l *Lexer) readChar() {
	// note: nextPos is initialised to 0
	// (why not just use l.position? kiv)
	if l.nextPos >= len(l.input) {
		l.c = 0 // set to null byte
	} else {
		l.c = l.input[l.nextPos]
	}
	l.position = l.nextPos
	l.nextPos += 1
}

func newToken(c byte, t token.TokenType) token.Token {
	if c == 0 {
		return token.Token{Type: t, Literal: ""}
	}
	return token.Token{Type: t, Literal: string(c)}
}

func isLetter(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || c == '_'
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

// func (l *Lexer) skipWhitespace() {
// 	for unicode.IsSpace(rune(l.c)) {
// 		l.readChar()
// 	}
// }

func (l *Lexer) NextToken() token.Token {
	// l.skipWhitespace()
	for unicode.IsSpace(rune(l.c)) {
		l.readChar()
	}

	var tok token.Token

	// note: we cannot just read single characters, because some seeking
	// (and internal state) will be needed
	switch l.c {

	// symbols
	case '=':
		tok = newToken(l.c, token.ASSIGN)
	case '+':
		tok = newToken(l.c, token.PLUS)
	case '(':
		tok = newToken(l.c, token.LPAREN)
	case ')':
		tok = newToken(l.c, token.RPAREN)
	case '{':
		tok = newToken(l.c, token.LBRACE)
	case '}':
		tok = newToken(l.c, token.RBRACE)
	case ',':
		tok = newToken(l.c, token.COMMA)
	case ';':
		tok = newToken(l.c, token.SEMICOLON)
	case 0:
		tok = newToken(l.c, token.EOF)

	default:
		if isLetter(l.c) {
			start := l.position
			for isLetter(l.c) {
				l.readChar()
			}
			tok.Literal = l.input[start:l.position]
			tok.SetType()
			return tok // don't call readChar again
		} else if isDigit(l.c) {
			start := l.position
			for isDigit(l.c) {
				l.readChar()
			}
			tok.Literal = l.input[start:l.position]
			tok.Type = token.INT
			return tok // don't call readChar again
		} else {
			tok = newToken(l.c, token.ILLEGAL)
			// panic(fmt.Sprintf("unhandled: '%s'", string(l.c)))
		}

	}
	l.readChar()
	return tok
}
