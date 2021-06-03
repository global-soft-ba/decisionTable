package data

//TODO: We have to care about, "Name" and VariableName , e.g. Name ="4 Eyes Principle", VariableName="FourEyesPrinciple"
//TODO: Extend to arbitrary navigation paths on structs, so far we only allow "Name.Key"

type Field struct {
	Name string
	Key  string
	Typ  DataTyp
}
