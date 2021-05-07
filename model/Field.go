package model

//TODO We have to care about, "Name" and VariableName , e.g. Name ="4 Eyes Principle", VariableName="4EyesPrinciple"

type Field struct {
	Name string
	Key  string
	Typ  VariableTyp
}
