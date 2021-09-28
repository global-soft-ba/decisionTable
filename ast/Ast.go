package ast

// The base Node interface which represents abstract syntac trees
type Node interface {
	String() string
	GetChildren() []Node
}
