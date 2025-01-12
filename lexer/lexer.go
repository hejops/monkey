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
	input    string // the human-readable string representation of a program
	position int    // current position
	c        byte   // character at current position
	nextPos  int    // for peeking; generally position+1
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // init
	return l
}

// readChar reads the current character and stores it in c. Then it advances
// the cursor.
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

// peekChar reads the next character without advancing the cursor.
func (l *Lexer) peekChar() byte {
	if l.nextPos >= len(l.input) {
		return 0
	} else {
		return l.input[l.nextPos]
	}
}

func isLetter(c byte) bool {
	return 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' || c == '_'
}

func isDigit(c byte) bool {
	return '0' <= c && c <= '9'
}

// NextToken reads tokens until an EOF is reached. It does not check for
// correctness of a program; that is the parser's job.
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

	// special cases:
	// = can be either = (ASSIGN) or == (EQUAL)
	// ! can be either ! (NOT) or != (NOT_EQUAL)

	if l.c == '=' && l.peekChar() == '=' {
		tok = token.Token{Type: token.EQUAL, Literal: "=="}
		l.readChar()
		l.readChar()

	} else if l.c == '!' && l.peekChar() == '=' {
		tok = token.Token{Type: token.NOT_EQUAL, Literal: "!="}
		l.readChar()
		l.readChar()

	} else if token.IsSymbol(l.c) {
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
