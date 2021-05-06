package termconverter

import "decisionTable/converters/grule/grlmodel"

type TermConverterInterface interface {
	ConvertExpression(expr grlmodel.Term) grlmodel.Term
	ConvertAssignments(expr grlmodel.Term) grlmodel.Term
}
