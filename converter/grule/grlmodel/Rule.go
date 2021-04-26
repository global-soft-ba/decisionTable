package grlmodel

type Rule struct {
	Name        string
	Description string
	Salience    int
	Expressions []Expression
	Assignments []Expression
}
