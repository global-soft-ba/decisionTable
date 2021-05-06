package expressionlanguages

import "decisionTable/converter/grule/grlmodel"

type ExpressionConverterInterface interface {
	ConvertExpression(expr grlmodel.Term) grlmodel.Term
	ConvertAssignments(expr grlmodel.Term) grlmodel.Term
}
