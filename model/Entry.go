package model

func CreateEntry(exp ExpressionInterface, exprLang ExpressionLanguage) Entry {
	return Entry{expression: exp, exprLang: exprLang}
}

type Entry struct {
	expression ExpressionInterface
	exprLang   ExpressionLanguage
}

func (e Entry) Expression() string {
	return e.expression.String()
}

func (e Entry) ExpressionLanguage() ExpressionLanguage {
	return e.exprLang
}
