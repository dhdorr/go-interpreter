package repl

import (
	"fmt"
	"io"

	"go-interpreter/evaluator"
	"go-interpreter/lexer"
	"go-interpreter/object"
	"go-interpreter/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer, test_map *object.Environment, session_id, code string) string {
	//scanner := bufio.NewScanner(in)

	// TODO whenever a user logs in, check sqlite to see if they exist, if not then make them
	// on login, add user to a dictionary of user(string:session code) -> environment
	// client must send their session code on every PUT request to maintain environment mapping
	// may have to move this users map to the main func and pass it in so it is not remade on every http request handle
	// limit history to like 10 successful inputs
	// test_map := make(map[string]*object.Environment)
	// env := object.NewEnvironment()
	// test_map["test"] = env

	//for {
	//fmt.Fprint(out, PROMPT)
	//scanned := scanner.Scan()
	//if !scanned {
	//	return
	//}

	//line := scanner.Text()
	// line := "2 + 2"
	l := lexer.New(code)

	// print lexical tokens
	// for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
	// 	fmt.Fprintf(out, "%+v\n", tok)
	// }

	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParserErrors(out, p.Errors())
		//continue
		return "errors"
	}

	evaluated := evaluator.Eval(program, test_map)
	result := "\n"
	if evaluated != nil {
		result = evaluated.Inspect()
		fmt.Printf(">> %s\n", result)
		// io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")
	}

	// io.WriteString(out, program.String())
	// io.WriteString(out, "\n")
	//}
	return result
}

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
