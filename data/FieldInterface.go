package data

type FieldInterface interface {
	Id() string
	DataTyp() DataTyp
	String() string
	GetQualifiedName() string
}
