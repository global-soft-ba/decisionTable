package validator

type DTableGruleExpressionValidator struct {
	valid  bool
	errors []error
}

func (d DTableGruleExpressionValidator) Validate() (bool, []error) {
	panic("implement me")
}
