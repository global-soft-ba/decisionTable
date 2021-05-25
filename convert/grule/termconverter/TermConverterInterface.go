package termconverter

import "decisionTable/convert/grule/grlmodel"

type TermConverterInterface interface {
	ConvertExpression(expr grlmodel.Term) grlmodel.Term
	ConvertAssignments(expr grlmodel.Term) grlmodel.Term
}
