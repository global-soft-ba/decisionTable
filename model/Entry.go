package model

const EmptyEntryExpressionToken = `-`

//TODO We have to care about, that assignment in DTable is just a value but in the ruleEngine an Set-Operator (like "=")

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
