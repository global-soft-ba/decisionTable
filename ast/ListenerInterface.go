package ast

type ListenerInterface interface {
	EnterNode(Node)
	ExitNode(Node)
}
