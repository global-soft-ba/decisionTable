package valid

type ValidatorInterface interface {
	Validate() (bool, []error)
	ValidateContainsInterferences() bool
}
