package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/jarsen/hax/evaluator"
	"github.com/jarsen/hax/lexer"
	"github.com/jarsen/hax/parser"
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
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "ERROR! DOES NOT COMPUTE!\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
