package data

import "decisionTable/data"

type RuleSet struct {
	Key             string
	Name            string
	HitPolicy       data.HitPolicy
	CollectOperator data.CollectOperator
	Interference    bool
	Rules           []Rule
}
