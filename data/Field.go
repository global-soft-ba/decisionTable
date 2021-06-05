package data

// This is just a example implementation of the FieldInterface

type Field struct {
	Name string
	Key  string
	Typ  DataTyp
}

func (f Field) Id() string {
	return f.Key + "." + f.Name
}

func (f Field) DataTyp() DataTyp {
	return f.Typ
}

func (f Field) String() string {
	return f.Id()
}
func (f Field) GetQualifiedName(separator string) string {
	return f.Name + separator + f.Key
}
