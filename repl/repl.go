package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/juandspy/monkey-lang/lexer"
	"github.com/juandspy/monkey-lang/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Print(PROMPT)
		// read the user input
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()

		l := lexer.New(line) // start a lexer with the user input
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			// print all the tokens we found
			fmt.Printf("%+v\n", tok)
		}
	}
}
