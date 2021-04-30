package model

func CreateEntry(exp string, exprLang ExpressionLanguage) Entry {
	return Entry{expression: exp, expressionLanguage: exprLang}
}

type Entry struct {
	expression         string
	expressionLanguage ExpressionLanguage
}

func (e Entry) Expression() string {
	return e.expression
}

func (e Entry) ExpressionLanguage() ExpressionLanguage {
	return e.expressionLanguage
}
