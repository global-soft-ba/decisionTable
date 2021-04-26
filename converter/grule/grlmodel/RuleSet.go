package grlmodel

import "decisionTable/model"

type RuleSet struct {
	Key             string
	Name            string
	HitPolicy       model.HitPolicy
	CollectOperator model.CollectOperator
	Rules           []Rule
}
