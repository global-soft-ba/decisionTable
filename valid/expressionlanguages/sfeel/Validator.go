package sfeel

import (
	"github.com/global-soft-ba/decisionTable/model"
	"github.com/global-soft-ba/decisionTable/parser/sfeel/parser"
)

func CreateValidator() Validator {
	return Validator{}
}

type Validator struct {
}

//TODO not sure if the antlr parser is thread safe. Until that, we use a new single parser instance for each validation

func (v Validator) ValidateInputEntry(field model.Field, entry model.Entry) (bool, []error) {
	p := parser.CreateSfeelParser(entry.Expression())

	switch field.Typ {
	case model.String:
		p.Parse().ValidStringInput()
	case model.Integer:
		p.Parse().ValidIntegerInput()
	case model.Double:
		p.Parse().ValidIntegerInput()
	case model.Boolean:
		p.Parse().ValidBoolInput()
	case model.Float:
		p.Parse().ValidNumberInput()
	case model.Long:
		p.Parse().ValidNumberInput()
	case model.DateTime:
		p.Parse().ValidDateTimeInput()
	}

	if len(p.Errors()) > 0 {
		return false, p.Errors()
	}

	return true, nil
}

func (v Validator) ValidateOutputEntry(field model.Field, entry model.Entry) (bool, []error) {
	p := parser.CreateSfeelParser(entry.Expression())

	switch field.Typ {
	case model.String:
		p.Parse().ValidStringOutput()
	case model.Integer:
		p.Parse().ValidIntegerOutput()
	case model.Double:
		p.Parse().ValidIntegerOutput()
	case model.Boolean:
		p.Parse().ValidBoolOutput()
	case model.Float:
		p.Parse().ValidNumberOutput()
	case model.Long:
		p.Parse().ValidNumberOutput()
	case model.DateTime:
		p.Parse().ValidDateTimeOutput()
	}

	if len(p.Errors()) > 0 {
		return false, p.Errors()
	}

	return true, nil
}
