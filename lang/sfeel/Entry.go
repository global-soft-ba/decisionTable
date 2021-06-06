package sfeel

import (
	"decisionTable/data"
	"decisionTable/lang/sfeel/antlr"
	sfeel "decisionTable/lang/sfeel/ast"
	"decisionTable/lang/sfeel/eval"
	"errors"
	"fmt"
	"reflect"
)

func CreateInputEntry(exp string) data.EntryInterface {
	tree, err := antlr.CreateParser(exp).Parse()
	evl := eval.CreateInputEntryEvaluator()
	if err != nil {
		return Entry{ast: nil, evaluator: evl, expression: exp}
	}
	return Entry{ast: tree, evaluator: evl, expression: exp}
}

func CreateOutputEntry(exp string) data.EntryInterface {
	tree, err := antlr.CreateParser(exp).Parse()
	evl := eval.CreateOutputEntryEvaluator()
	if err != nil {
		return Entry{ast: nil, evaluator: evl, expression: exp}
	}
	return Entry{ast: tree, evaluator: evl, expression: exp}
}

type Entry struct {
	ast        sfeel.Node
	evaluator  eval.EvaluatorInterface
	expression string
}

func (e Entry) String() string {
	return e.expression
}
func (e Entry) ExpressionLanguage() data.ExpressionLanguage {
	return data.SFEEL
}

func (e Entry) Validate() (bool, []error) {
	if e.ast == nil {
		_, err := antlr.CreateParser(e.expression).Parse()
		return false, err
	}
	return e.evaluator.Eval(e.ast)
}
func (e Entry) ValidateDataTypeOfExpression(varType data.DataTyp) (bool, error) {
	if e.ast == nil {
		return false, nil
	}

	switch e.ast.GetOperandDataType() {
	case reflect.TypeOf(sfeel.EmptyStatement{}):
		return true, nil
	case reflect.TypeOf(sfeel.QualifiedName{}):
		return true, nil
	case reflect.TypeOf(sfeel.Integer{}):
		if varType == data.Integer {
			return true, nil
		}
	case reflect.TypeOf(sfeel.Float{}):
		if varType == data.Float {
			return true, nil
		}
	case reflect.TypeOf(sfeel.String{}):
		if varType == data.String {
			return true, nil
		}
	case reflect.TypeOf(sfeel.Boolean{}):
		if varType == data.Boolean {
			return true, nil
		}
	case reflect.TypeOf(sfeel.DateTime{}):
		if varType == data.DateTime {
			return true, nil
		}
	}

	return false, errors.New(fmt.Sprintf("given data type %s is not compatible with %s", varType, e.ast.GetOperandDataType()))
}
func (e Entry) ValidateExistenceOfFieldReferencesInExpression(fields []data.FieldInterface) ([]data.FieldInterface, []error) {
	qualifiedFields := sfeel.GetAllQualifiedNames(e.ast)
	var errOut []error
	var out []data.FieldInterface

	for _, val := range qualifiedFields {
		qf, err := e.getFieldUsingQualifiedName(val, fields)
		if err != nil {
			errOut = append(errOut, err)
		} else {
			out = append(out, qf)
		}
	}

	return out, errOut
}

func (e Entry) getFieldUsingQualifiedName(name sfeel.QualifiedName, fields []data.FieldInterface) (data.FieldInterface, error) {
	for _, val := range fields {
		if val.GetQualifiedName() == name.GetQualifiedName() {
			return val, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("couldn't find qualified name %s in field list", name.String()))
}

func (e Entry) Convert(listener sfeel.SFeelListenerInterface) {
	tree := sfeel.CreateSFeelTreeWalker(listener)
	tree.Walk(e.ast)
}
