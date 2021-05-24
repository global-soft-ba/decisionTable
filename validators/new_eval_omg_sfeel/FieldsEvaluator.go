package new_eval_omg_sfeel

import (
	"decisionTable/SFeel/ast"
	"decisionTable/SFeel/errors"
	"decisionTable/model"
	"reflect"
)

type FieldsEvaluator struct {
	fields []model.Field
	field  model.Field
}

func (e FieldsEvaluator) Eval(curField model.Field, fields []model.Field) (bool, []error) {
	e.field = curField
	e.fields = fields

	return true, nil
}

func (e FieldsEvaluator) EvalTest(unaryTest ast.UnaryTest) (bool, error) {

	return true, nil
}

func (e FieldsEvaluator) evalDataType(data ast.Node) (bool, error) {
	switch data.(type) {
	case ast.Integer:
		if e.field.Typ == model.Integer {
			return true, nil
		}
	case ast.Float:
		if e.field.Typ == model.Float {
			return true, nil
		}
	case ast.String:
		if e.field.Typ == model.String {
			return true, nil
		}
	case ast.Boolean:
		if e.field.Typ == model.Boolean {
			return true, nil
		}
	case ast.DateTime:
		if e.field.Typ == model.DateTime {
			return true, nil
		}
	case ast.QualifiedName:
		return e.evalQualifiedName(data.(ast.QualifiedName))
	}

	return false, errors.NewError("field data type %s does not expected Sfeel data typ %s ", e.field.Typ, reflect.TypeOf(data))
}

func (e FieldsEvaluator) evalQualifiedName(name ast.QualifiedName) (bool, error) {
	res, err := e.evalQualifiedNameExistenceInFields(name)
	if err != nil {
		return false, err
	}
	if res.Typ != e.field.Typ {
		return false, errors.NewError("referenced field type %s doesn't correspond to field typ %s ", e.field.Typ, reflect.TypeOf(res.Typ))

	}
	return true, nil

}

func (e FieldsEvaluator) evalQualifiedNameExistenceInFields(name ast.QualifiedName) (model.Field, error) {
	for _, val := range e.fields {
		//TODO Extend to allow arbitrary navigation paths on structs
		if val.Name == name.Value[0] && val.Key == name.Value[1] {
			return val, nil
		}
	}
	return model.Field{}, errors.NewError("couldn't find qualified name %s in field list", name.String())
}
