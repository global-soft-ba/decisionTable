package conv

import (
	"decisionTable/data"
	"fmt"
)

func CreateSFeelConverter(listener SFeelListener) SFeelConverter {
	return SFeelConverter{listener: listener}
}

type SFeelConverter struct {
	listener SFeelListener
}

func (c SFeelConverter) Convert(e data.EntryInterface) {
	listener := e.Convert(&c.listener)
	fmt.Println("Listener=", listener)
}
