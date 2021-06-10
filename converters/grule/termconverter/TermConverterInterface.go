package termconverter

import "github.com/global-soft-ba/decisionTable/converters/grule/grlmodel"

type TermConverterInterface interface {
	ConvertExpression(expr grlmodel.Term) grlmodel.Term
	ConvertAssignments(expr grlmodel.Term) grlmodel.Term
}
