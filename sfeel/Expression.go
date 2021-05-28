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

func CreateInputExpression(exp string) (model.ExpressionInterface, []error) {
	tree, err := antlr.CreateParser(exp).Parse()
	if err != nil {
		return Expression{}, err
	}

	evl, err := eval.CreateInputEntryEvaluator().Eval(tree)
	if !evl {
		return Expression{}, err
	}

	return Expression{ast: tree, expression: exp}, nil
}
func CreateOutputExpression(exp string) (model.ExpressionInterface, []error) {
	tree, err := antlr.CreateParser(exp).Parse()
	if err != nil {
		return Expression{}, err
	}

	evl, err := eval.CreateOutputEntryEvaluator().Eval(tree)
	if !evl {
		return Expression{}, err
	}

	return Expression{ast: tree, expression: exp}, nil
}

type Expression struct {
	ast        ast.Node
	expression string
}

func (e Expression) String() string {
	if e.ast == nil {
		return ""
	}
	return e.ast.String()
}
func (e Expression) ValidateDataTypeOfExpression(varType model.DataTyp) (bool, error) {
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
func (e Expression) ValidateExistenceOfFieldReferencesInExpression(fields []model.Field) ([]model.Field, []error) {
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
func (e Expression) getFieldUsingQualifiedName(name ast.QualifiedName, fields []model.Field) (model.Field, error) {
	for _, val := range fields {
		//TODO Extend to allow arbitrary navigation paths on structs
		if val.Name == name.Value[0] && val.Key == name.Value[1] {
			return val, nil
		}
	}
	return model.Field{}, errors.New(fmt.Sprintf("couldn't find qualified name %s in field list", name.String()))
}
