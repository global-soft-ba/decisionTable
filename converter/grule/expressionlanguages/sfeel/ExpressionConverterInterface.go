package sfeel

import (
	parser2 "decisionTable/parser/sfeel/parser"
)

type ExpressionConverterInterface interface {
	Convert(expression string, parser parser2.Parser) string
}
