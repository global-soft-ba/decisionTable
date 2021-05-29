package sfeel

import (
	"decisionTable/model"
	"decisionTable/sfeel/antlr"
	"decisionTable/sfeel/ast"
	"decisionTable/sfeel/eval"
	"errors"
	"fmt"
	"reflect"
)

func CreateInputEntry(exp string) model.EntryInterface {
	tree, err := antlr.CreateParser(exp).Parse()
	evl := eval.CreateInputEntryEvaluator()
	if err != nil {
		return Entry{ast: nil, evaluator: evl, expression: exp}
	}
	return Entry{ast: tree, evaluator: evl, expression: exp}
}

func CreateOutputEntry(exp string) model.EntryInterface {
	tree, err := antlr.CreateParser(exp).Parse()
	evl := eval.CreateOutputEntryEvaluator()
	if err != nil {
		return Entry{ast: nil, evaluator: evl, expression: exp}
	}
	return Entry{ast: tree, evaluator: evl, expression: exp}
}

type Entry struct {
	ast        ast.Node
	evaluator  eval.EvaluatorInterface
	expression string
}

func (e Entry) String() string {
	return e.expression
}
func (e Entry) ExpressionLanguage() model.ExpressionLanguage {
	return model.SFEEL
}

func (e Entry) Validate() (bool, []error) {
	if e.ast == nil {
		_, err := antlr.CreateParser(e.expression).Parse()
		return false, err
	}
	return e.evaluator.Eval(e.ast)
}
func (e Entry) ValidateDataTypeOfExpression(varType model.DataTyp) (bool, error) {
	if e.ast == nil {
		return false, nil
	}

	switch e.ast.GetOperandDataType() {
	case reflect.TypeOf(ast.EmptyStatement{}):
		return true, nil
	case reflect.TypeOf(ast.QualifiedName{}):
		return true, nil
	case reflect.TypeOf(ast.Integer{}):
		if varType == model.Integer {
			return true, nil
		}
	case reflect.TypeOf(ast.Float{}):
		if varType == model.Float {
			return true, nil
		}
	case reflect.TypeOf(ast.String{}):
		if varType == model.String {
			return true, nil
		}
	case reflect.TypeOf(ast.Boolean{}):
		if varType == model.Boolean {
			return true, nil
		}
	case reflect.TypeOf(ast.DateTime{}):
		if varType == model.DateTime {
			return true, nil
		}
	}

	return false, errors.New(fmt.Sprintf("given data type %s is not compatible with %s", varType, e.ast.GetOperandDataType()))
}
func (e Entry) ValidateExistenceOfFieldReferencesInExpression(fields []model.Field) ([]model.Field, []error) {
	qualifiedFields := ast.GetAllQualifiedNames(e.ast)
	var errOut []error
	var out []model.Field

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

func (e Entry) getFieldUsingQualifiedName(name ast.QualifiedName, fields []model.Field) (model.Field, error) {
	for _, val := range fields {
		//TODO Extend to allow arbitrary navigation paths on structs
		if val.Name == name.Value[0] && val.Key == name.Value[1] {
			return val, nil
		}
	}
	return model.Field{}, errors.New(fmt.Sprintf("couldn't find qualified name %s in field list", name.String()))
}
