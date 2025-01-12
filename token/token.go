package token

const (
	// when defining a new token, make sure to include it in the
	// corresponding map

	// ILLEGAL   TokenType = "ILLEGAL"
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	IDENT   = "IDENT"
	INT     = "INT"

	// keywords

	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	RETURN   = "RETURN"

	// symbols (these are strings, to be consistent with the above)

	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	MULTIPLY  = "*"
	DIVIDE    = "/"
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LESS      = "<"
	GREATER   = ">"
	NOT       = "!"
	EQUAL     = "=="
	NOT_EQUAL = "!="
)

type (
	TokenType string // should be enum of strings?

	Token struct {
		Type    TokenType
		Literal string // the Token's literal text representation
	}
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
}

var symbols = map[byte]TokenType{
	'=': ASSIGN,
	'+': PLUS,
	'-': MINUS,
	'*': MULTIPLY,
	'/': DIVIDE,
	',': COMMA,
	';': SEMICOLON,
	'(': LPAREN,
	')': RPAREN,
	'{': LBRACE,
	'}': RBRACE,
	'<': LESS,
	'>': GREATER,
	'!': NOT,
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
