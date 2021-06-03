package termconverter

import (
	grlmodel2 "decisionTable/conv/grule/data"
)

type TermConverterInterface interface {
	ConvertExpression(expr grlmodel2.Term) grlmodel2.Term
	ConvertAssignments(expr grlmodel2.Term) grlmodel2.Term
}
