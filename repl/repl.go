package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/jarsen/hax/lexer"
	"github.com/jarsen/hax/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	s := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := s.Scan()
		if !scanned {
			return
		}

		line := s.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%v\n", tok)
		}
	}
}
