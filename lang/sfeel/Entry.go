package sfeel

import (
	"decisionTable/data"
	antlr2 "decisionTable/lang/sfeel/antlr"
	ast "decisionTable/lang/sfeel/ast"
	"decisionTable/lang/sfeel/conv"
	eval2 "decisionTable/lang/sfeel/eval"
	"errors"
	"fmt"
	"reflect"
)

func CreateInputEntry(exp string) data.EntryInterface {
	tree, err := antlr2.CreateParser(exp).Parse()
	evl := eval2.CreateInputEntryEvaluator()
	if err != nil {
		return Entry{ast: nil, evaluator: evl, expression: exp}
	}
	return Entry{ast: tree, evaluator: evl, expression: exp}
}

func CreateOutputEntry(exp string) data.EntryInterface {
	tree, err := antlr2.CreateParser(exp).Parse()
	evl := eval2.CreateOutputEntryEvaluator()
	if err != nil {
		return Entry{ast: nil, evaluator: evl, expression: exp}
	}
	return Entry{ast: tree, evaluator: evl, expression: exp}
}

type Entry struct {
	ast        ast.Node
	evaluator  eval2.EvaluatorInterface
	expression string
}

//Use within decisionTable
func (e Entry) String() string {
	return e.expression
}
func (e Entry) ExpressionLanguage() data.ExpressionLanguage {
	return data.SFEEL
}

func (e Entry) Validate() (bool, []error) {
	if e.ast == nil {
		_, err := antlr2.CreateParser(e.expression).Parse()
		return false, err
	}
	return e.evaluator.Eval(e.ast)
}
func (e Entry) ValidateDataTypeOfExpression(varType data.DataTyp) (bool, error) {
	if e.ast == nil {
		return false, nil
	}

	switch e.ast.GetOperandDataType() {
	case reflect.TypeOf(ast.EmptyStatement{}):
		return true, nil
	case reflect.TypeOf(ast.QualifiedName{}):
		return true, nil
	case reflect.TypeOf(ast.Integer{}):
		if varType == data.Integer {
			return true, nil
		}
	case reflect.TypeOf(ast.Float{}):
		if varType == data.Float {
			return true, nil
		}
	case reflect.TypeOf(ast.String{}):
		if varType == data.String {
			return true, nil
		}
	case reflect.TypeOf(ast.Boolean{}):
		if varType == data.Boolean {
			return true, nil
		}
	case reflect.TypeOf(ast.DateTime{}):
		if varType == data.DateTime {
			return true, nil
		}
	}

	return false, errors.New(fmt.Sprintf("given data type %s is not compatible with %s", varType, e.ast.GetOperandDataType()))
}
func (e Entry) ValidateExistenceOfFieldReferencesInExpression(fields []data.Field) ([]data.Field, []error) {
	qualifiedFields := ast.GetAllQualifiedNames(e.ast)
	var errOut []error
	var out []data.Field

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

func (e Entry) getFieldUsingQualifiedName(name ast.QualifiedName, fields []data.Field) (data.Field, error) {
	for _, val := range fields {
		//TODO Extend to allow arbitrary navigation paths on structs
		if val.Name == name.Value[0] && val.Key == name.Value[1] {
			return val, nil
		}
	}
	return data.Field{}, errors.New(fmt.Sprintf("couldn't find qualified name %s in field list", name.String()))
}

func (e Entry) Convert(listener conv.SFeelBaseListenerInterface) conv.SFeelBaseListenerInterface {
	tree := conv.CreateTreeWalker(listener)
	return tree.Walk(e.ast)
}
