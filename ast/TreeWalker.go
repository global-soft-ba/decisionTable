package ast

func CreateTreeWalker(listener ListenerInterface) TreeWalker {
	return TreeWalker{listener: listener}
}

type TreeWalker struct {
	listener ListenerInterface
}

func (w TreeWalker) Walk(node Node) {
	w.listener.EnterNode(node)

	//Children of Node
	for _, val := range node.GetChildren() {
		w.Walk(val)
	}

	//Node
	w.listener.ExitNode(node)
}
