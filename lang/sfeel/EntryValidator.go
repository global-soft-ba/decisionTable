package sfeel

import (
	"errors"
	"fmt"
	"github.com/global-soft-ba/decisionTable/data"
	"github.com/global-soft-ba/decisionTable/data/dataType"
	"github.com/global-soft-ba/decisionTable/data/field"
	"github.com/global-soft-ba/decisionTable/lang/sfeel/antlr"
	sfeel "github.com/global-soft-ba/decisionTable/lang/sfeel/ast"
	"github.com/global-soft-ba/decisionTable/lang/sfeel/eval"
	"reflect"
)

var (
	ErrInvalidAst = errors.New("invalid ast, can not be null in Validate() method")
)

type EntryValidator struct {
	ast       sfeel.Node
	evaluator eval.EvaluatorInterface
}

func CreateInputEntryValidator(entry string) data.EntryValidatorInterface {
	tree, err := antlr.CreateParser(entry).ParseInput()
	evl := eval.CreateInputEntryEvaluator()
	if err != nil {
		return EntryValidator{ast: nil, evaluator: evl}
	}

	return EntryValidator{ast: tree, evaluator: evl}
}

func CreateOutputEntryValidator(entry string) data.EntryValidatorInterface {
	tree, err := antlr.CreateParser(entry).ParseOutput()
	evl := eval.CreateOutputEntryEvaluator()
	if err != nil {
		return EntryValidator{ast: nil, evaluator: evl}
	}

	return EntryValidator{ast: tree, evaluator: evl}
}

func (e EntryValidator) Validate() (bool, []error) {
	if e.ast == nil {
		return false, []error{ErrInvalidAst}
	}

	return e.evaluator.Eval(e.ast)
}

func (e EntryValidator) ValidateDataTypeOfExpression(varType dataType.DataType) (bool, error) {
	if e.ast == nil {
		return false, nil
	}

	switch e.ast.GetOperandDataType() {
	case reflect.TypeOf(sfeel.EmptyStatement{}):
		return true, nil
	case reflect.TypeOf(sfeel.QualifiedName{}):
		return true, nil
	case reflect.TypeOf(sfeel.Integer{}):
		if varType == dataType.Integer {
			return true, nil
		}
	case reflect.TypeOf(sfeel.Float{}):
		if varType == dataType.Float {
			return true, nil
		}
	case reflect.TypeOf(sfeel.String{}):
		if varType == dataType.String {
			return true, nil
		}
	case reflect.TypeOf(sfeel.Boolean{}):
		if varType == dataType.Boolean {
			return true, nil
		}
	case reflect.TypeOf(sfeel.DateTime{}):
		if varType == dataType.DateTime {
			return true, nil
		}
	case reflect.TypeOf(sfeel.UnaryTests{}):
		return true, nil
	case reflect.TypeOf(sfeel.UnaryTest{}):
		return true, nil
	case reflect.TypeOf(sfeel.Interval{}):
		return true, nil
	case reflect.TypeOf(sfeel.Parentheses{}):
		return true, nil
	case reflect.TypeOf(sfeel.SimpleExpression{}):
		return true, nil
	case reflect.TypeOf(sfeel.SimpleValue{}):
		return true, nil
	}

	return false, errors.New(fmt.Sprintf("given data type %s is not compatible with %s", varType, e.ast.GetOperandDataType()))
}

func (e EntryValidator) ValidateExistenceOfFieldReferencesInExpression(fields []field.Field) ([]field.Field, []error) {
	qualifiedFields := sfeel.GetAllQualifiedNames(e.ast)
	var errOut []error
	var out []field.Field
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

func (e EntryValidator) getFieldUsingQualifiedName(name sfeel.QualifiedName, fields []field.Field) (field.Field, error) {
	for _, f := range fields {
		if f.Name == name.GetQualifiedName() {
			return f, nil
		}
	}
	return field.Field{}, errors.New(fmt.Sprintf("couldn't find qualified name %s in field list", name.String()))
}
