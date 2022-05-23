package data

import (
	"github.com/global-soft-ba/decisionTable/data/collectOperator"
	"github.com/global-soft-ba/decisionTable/data/hitPolicy"
)

type RuleSet struct {
	Key             string
	Name            string
	HitPolicy       hitPolicy.HitPolicy
	CollectOperator collectOperator.CollectOperator
	Interference    bool
	Rules           []Rule
}
