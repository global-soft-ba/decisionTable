package grlmodel

import "github.com/global-soft-ba/decisionTable/model"

type RuleSet struct {
	Key             string
	Name            string
	HitPolicy       model.HitPolicy
	CollectOperator model.CollectOperator
	Interference    bool
	Rules           []Rule
}
