package main

import (
	"bufio"
	"bytes"
	"decisionTable/sfeel/antlr"
	parsererror "decisionTable/sfeel/antlr/errors"
	"decisionTable/sfeel/ast"
	"decisionTable/sfeel/eval"
	"fmt"
	"io"
	"os"
	"strconv"
)

const PROMPT = ">> "
const REPLY = "==="

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		exp := scanner.Text()

		tree, err := antlr.CreateParser(exp).Parse()
		if len(err) != 0 {
			printParserErrors(out, err)
			continue
		}
		io.WriteString(out, REPLY+"identified expression as type: "+printTreeTypes(ast.GetAllTreeNodeTypes(tree))+"\n")

		_, err = eval.CreateInputEntryEvaluator().Eval(tree)
		if len(err) != 0 {
			printEvaluatorErrors(out, err, "input")
		} else {
			io.WriteString(out, REPLY+"input evaluator: VALID"+"\n")
		}

		_, err = eval.CreateOutputEntryEvaluator().Eval(tree)
		if len(err) != 0 {
			printEvaluatorErrors(out, err, "output")
		} else {
			io.WriteString(out, REPLY+"output evaluator: VALID"+"\n")
		}

		io.WriteString(out, "\n")

	}
}

func printTreeTypes(tree []string) string {
	var out bytes.Buffer
	for i := len(tree) - 1; i >= 0; i-- {
		out.WriteString(tree[i])
		if i != 0 {
			out.WriteString("->")
		}
	}

	return out.String()
}

func printEvaluatorErrors(out io.Writer, errors []error, env string) {
	for _, msg := range errors {
		io.WriteString(out, REPLY+env+" evaluator: "+msg.Error()+"\n")
	}
}

func printParserErrors(out io.Writer, errors []error) {
	io.WriteString(out, REPLY+"parser errors:\n")
	for _, msg := range errors {
		err := msg.(parsererror.ExpressionSyntaxError)
		line := strconv.Itoa(err.Line())
		col := strconv.Itoa(err.Column())

		io.WriteString(out, "   "+line+":"+col+" "+err.Msg()+"\n")
	}
}

func main() {

	fmt.Printf("This is the SFEEL expression language!\n")
	fmt.Printf("Feel free to type in expressions\n")
	Start(os.Stdin, os.Stdout)

}
