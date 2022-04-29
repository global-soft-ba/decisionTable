package data

// This is just an example implementation of the FieldInterface used in test cases

type TestField struct {
	Name string
	Key  string
	Type DataType
}

func (f TestField) ID() string {
	if f.Key == "" && f.Name == "" {
		return ""
	}

	return f.Key + "." + f.Name
}

func (f TestField) DataType() DataType {
	return f.Type
}

func (f TestField) String() string {
	return f.ID()
}
func (f TestField) GetQualifiedName() string {
	return f.Name + "." + f.Key
}
