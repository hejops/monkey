package token

const (
	// ILLEGAL   TokenType = "ILLEGAL"
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	IDENT     = "IDENT"
	INT       = "INT"
	ASSIGN    = "="
	PLUS      = "+"
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	FUNCTION  = "FUNCTION"
	LET       = "LET"
)

type (
	TokenType string // should be enum of strings?

	Token struct {
		Type    TokenType
		Literal string // the Token's literal text representation
	}
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

var symbols = map[byte]TokenType{
	'=': ASSIGN,
	'+': PLUS,
	',': COMMA,
	';': SEMICOLON,
	'(': LPAREN,
	')': RPAREN,
	'{': LBRACE,
	'}': RBRACE,
	0:   EOF,
}

func IsSymbol(c byte) bool {
	_, ok := symbols[c]
	return ok
}

func NewSymbol(c byte) Token {
	t := symbols[c]
	if c == 0 {
		return Token{Type: t, Literal: ""}
	}
	return Token{Type: t, Literal: string(c)}
}

func (t *Token) SetType() {
	if kw, ok := keywords[t.Literal]; ok {
		t.Type = kw
		// } else if {
		// 	t.Type=INT
	} else {
		t.Type = IDENT
	}
}
