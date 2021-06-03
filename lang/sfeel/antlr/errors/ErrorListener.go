package errors

import (
	"errors"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func NewError(format string, a ...interface{}) error {
	return errors.New(fmt.Sprintf(format, a...))
}

type ErrorListener struct {
	*antlr.DefaultErrorListener // Embed default which ensures we fit the interface
	Errors                      []error
}

func (c *ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	c.Errors = append(c.Errors, ExpressionSyntaxError{line, column, msg})
}
