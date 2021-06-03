package conv

import (
	sfeel "decisionTable/lang/sfeel/ast"
	"decisionTable/lang/sfeel/conv"
	"strconv"
)

type SFeelListener struct {
	conv.SFeelBaseListener
	data string
}

func (l *SFeelListener) ExitInteger(ctx sfeel.Integer) {

	l.data = strconv.FormatInt(ctx.Value, 10)
}
