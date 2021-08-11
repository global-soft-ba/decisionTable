package data

// This is just a example implementation of the FieldInterface used in test cases

type TestField struct {
	Name string
	Key  string
	Typ  DataTyp
}

func (f TestField) Id() string {
	if f.Key == "" && f.Name == "" {
		return ""
	}

	return f.Key + "." + f.Name
}

func (f TestField) DataTyp() DataTyp {
	return f.Typ
}

func (f TestField) String() string {
	return f.Id()
}
func (f TestField) GetQualifiedName() string {
	return f.Name + "." + f.Key
}
