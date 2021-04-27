package grlmodel

type Rule struct {
	Name        string
	Description string
	Salience    int
	InvSalience int //Necessary for HitPolicies
	Expressions []Expression
	Assignments []Expression
}
