package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/parser"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			fmt.Println("Whoops! Parsing Errors:")
			for _, msg := range p.Errors() {
				fmt.Println("\t", msg)
			}
			continue
		}
		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}
