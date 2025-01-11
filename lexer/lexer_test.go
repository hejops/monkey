package lexer

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"monkey/token"
)

func TestSymbols(t *testing.T) {
	lexer := New(`=+(){},;`)

	chars := []struct {
		_type token.TokenType
		s     string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
	}

	for _, tt := range chars {
		tok := lexer.NextToken()
		assert.Equal(t, tok.Type, tt._type)
		assert.Equal(t, tok.Literal, tt.s)
	}
}

func TestSingleAssignment(t *testing.T) {
	lexer := New(`let five = 5;`)

	tokens := []struct {
		_type token.TokenType
		s     string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	for _, tt := range tokens {
		tok := lexer.NextToken()
		assert.Equal(t, tok.Type, tt._type)
		assert.Equal(t, tok.Literal, tt.s)
	}
}

// func TestBasicProgram(t *testing.T) {
// 	lexer := New(`let five = 5;
// let ten = 10;
// let add = fn(x, y) {
// 	x + y;
// };
// let result = add(five, ten);
// `)
//
// 	tokens := []struct {
// 		_type token.TokenType
// 		s     string
// 	}{
// 		{token.LET, "="},
// 		{token.IDENT, "five"},
// 		{token.ASSIGN, "="},
// 		{token.INT, "5"},
// 	}
//
// 	for _, tt := range tokens {
// 		tok := lexer.NextToken()
// 		assert.Equal(t, tok.Type, tt._type)
// 		assert.Equal(t, tok.Literal, tt.s)
// 	}
// }
