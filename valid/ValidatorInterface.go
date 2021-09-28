package valid

import "github.com/global-soft-ba/decisionTable/data"

type ValidatorInterface interface {
	Validate(table data.Table) (bool, []error)
	ValidateContainsInterferences(table data.Table) bool
}
