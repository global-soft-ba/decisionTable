package model

func CreateEntry(exp string, exprLang ExpressionLanguage) Entry {
	return Entry{expression: exp, expressionLanguage: exprLang, emptyExpression: false}
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
