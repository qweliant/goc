package repl

import (
	"bufio"
	"fmt"
	"goc/token"
	"goc/lexer"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text() // scanner.Text() returns the current input, line of text
		l := lexer.New(line)   // lexer.New() returns a pointer to a lexer struct

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok) // %+v adds field names
		}
	}
}
