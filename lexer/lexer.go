package lexer

import (
	"fmt"
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

func isLetter(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || c == '_'
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

func (l *Lexer) NextToken() token.Token {
	for unicode.IsSpace(rune(l.c)) {
		l.readChar()
	}

	// there are only five valid possibilities:
	// whitespace (ignored)
	// symbol
	// string - identifier (variables and functions; the distinction is not important here)
	// string - keyword
	// int

	var tok token.Token

	if token.IsSymbol(l.c) {
		tok = token.NewSymbol(l.c)
		l.readChar()

	} else if isLetter(l.c) {
		start := l.position
		for isLetter(l.c) {
			l.readChar()
		}
		tok.Literal = l.input[start:l.position]
		tok.SetType()

	} else if isDigit(l.c) {
		start := l.position
		for isDigit(l.c) {
			l.readChar()
		}
		tok.Literal = l.input[start:l.position]
		tok.Type = token.INT

	} else {
		// tok = token.Token{Type: token.ILLEGAL,Literal: string(l.c)}
		// l.readChar()
		panic(fmt.Sprintf("unhandled: '%s'", string(l.c)))
	}

	return tok
}
