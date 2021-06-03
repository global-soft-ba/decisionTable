package conv

import (
	"decisionTable/lang/sfeel/conv"
	"fmt"
)

type SFeelListener struct {
	conv.SFeelBaseListener
	data string
}

func (l *SFeelListener) ExitInteger() {
	fmt.Println("Hallo Integer")
	l.data = "Hello Integer"
}
