package data

type FieldInterface interface {
	ID() string
	DataType() DataType
	String() string
	GetQualifiedName() string
}
