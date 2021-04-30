package validator

type ValidatorInterface interface {
	Validate() (bool, []error)
	ValidateInterferences() bool
}
