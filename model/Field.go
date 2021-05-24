package model

//TODO: We have to care about, "Name" and VariableName , e.g. Name ="4 Eyes Principle", VariableName="FourEyesPrinciple"
// - We should use parser rule qualified_name to verify, if the name is valid in SFeel expressions because of qualified_name
//   references
//TODO: Extend to arbitrary navigation paths on structs, so far we only allow "Name.Key"

type Field struct {
	Name string
	Key  string
	Typ  VariableTyp
}
