package repl

import (
	"bufio"
	"fmt"
	"io"
	"orj/lexer"
	"orj/token"
)

const PROMPT = ">> "

func Start(input io.Reader, output io.Writer) {
	scanner := bufio.NewScanner(input)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		lexerObj := lexer.New(line)

		for tok := lexerObj.NextToken(); tok.Type != token.EOF; tok = lexerObj.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
