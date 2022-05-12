package data

type Rule struct {
	Name        string
	Annotation  string
	Salience    int
	InvSalience int //Necessary for HitPolicies
	Expressions []Term
	Assignments []Term
}
