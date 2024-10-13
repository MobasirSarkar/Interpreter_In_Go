package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/MobasirSarkar/Interpreter_In_Go/evaluator"
	"github.com/MobasirSarkar/Interpreter_In_Go/lexer"
	"github.com/MobasirSarkar/Interpreter_In_Go/object"
	"github.com/MobasirSarkar/Interpreter_In_Go/parser"
)

const PROMPT = "➜ "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

const MOBASIR_SARKAR = `
    /\      |\    /|  
   /__\     | \  / |  
  /    \    |  \/  |  
   `

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MOBASIR_SARKAR)
	io.WriteString(out, "Oops ! We ran into Some Problem!\n")
	io.WriteString(out, "parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
