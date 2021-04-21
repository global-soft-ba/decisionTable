package validator

type DTableValidatorInterface interface {
	Validate() (bool, []error)
}
