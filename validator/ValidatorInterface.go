package validator

type ValidatorInterface interface {
	Validate() (bool, []error)
	ValidateContainsInterferences() bool
}
