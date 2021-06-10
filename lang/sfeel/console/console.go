package main

import (
	"bufio"
	"bytes"
	"fmt"
	antlr2 "github.com/global-soft-ba/decisionTable/lang/sfeel/antlr"
	errors2 "github.com/global-soft-ba/decisionTable/lang/sfeel/antlr/errors"
	ast2 "github.com/global-soft-ba/decisionTable/lang/sfeel/ast"
	eval2 "github.com/global-soft-ba/decisionTable/lang/sfeel/eval"
	"io"
	"os"
	"strconv"
)

const PROMPT = ">> "
const REPLY = "==> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		exp := scanner.Text()

		tree, err := antlr2.CreateParser(exp).Parse()
		if len(err) != 0 {
			printParserErrors(out, err)
			continue
		}
		io.WriteString(out, REPLY+"parser literal: "+tree.ParserLiteral()+"\n")
		io.WriteString(out, REPLY+"ast string representation: "+tree.String()+"\n")
		io.WriteString(out, REPLY+"identified expression as type: "+printTreeTypes(ast2.GetAllTreeNodeTypes(tree))+"\n")

		_, err = eval2.CreateInputEntryEvaluator().Eval(tree)
		if len(err) != 0 {
			printEvaluatorErrors(out, err, "input")
		} else {
			io.WriteString(out, REPLY+"input evaluator: VALID"+"\n")
		}

		io.WriteString(out, REPLY+"semantic evaluator: NOT IMPLEMENTED"+"\n")

		_, err = eval2.CreateOutputEntryEvaluator().Eval(tree)
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
		err := msg.(errors2.ExpressionSyntaxError)
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
