package validators

type ValidatorInterface interface {
	Validate() (bool, []error)
	ValidateContainsInterferences() bool
}
