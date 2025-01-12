package repl

import (
	"bufio"
	"fmt"
	"os"

	"monkey/lexer"
	"monkey/token"
)

const prompt = "monkey> " // sqlite3

func Start() {
	sc := bufio.NewScanner(os.Stdin)
	out := os.Stdout

	fmt.Print(prompt)
	for sc.Scan() {
		line := sc.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}

		fmt.Print(prompt)
	}
}
