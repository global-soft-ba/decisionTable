package model

const EmptyEntryExpressionToken = `-`

func CreateEntry(exp string, exprLang ExpressionLanguage) Entry {
	entry := Entry{expression: exp, expressionLanguage: exprLang, emptyExpression: false}
	if exp == EmptyEntryExpressionToken {
		entry.emptyExpression = true
	}
	return entry
}

type Entry struct {
	expression         string
	expressionLanguage ExpressionLanguage
	emptyExpression    bool
}

func (e Entry) Expression() string {
	return e.expression
}

func (e Entry) ExpressionLanguage() ExpressionLanguage {
	return e.expressionLanguage
}

func (e Entry) EmptyExpression() bool {
	return e.emptyExpression
}
