package errors

import "fmt"

type ExpressionSyntaxError struct {
	line   int
	column int
	msg    string
}

func (c ExpressionSyntaxError) Line() int {
	return c.line
}

func (c ExpressionSyntaxError) Column() int {
	return c.column
}

func (c ExpressionSyntaxError) Msg() string {
	return c.msg
}

func (c ExpressionSyntaxError) Error() string {
	return fmt.Sprintf("%d:%d: syntax error:%v", c.line, c.column, c.msg)
}
