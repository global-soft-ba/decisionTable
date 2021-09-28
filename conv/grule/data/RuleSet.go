package data

import "github.com/global-soft-ba/decisionTable/data"

type RuleSet struct {
	Key             string
	Name            string
	HitPolicy       data.HitPolicy
	CollectOperator data.CollectOperator
	Interference    bool
	Rules           []Rule
}
