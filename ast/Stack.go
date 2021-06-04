package ast

type (
	Stack struct {
		top    *node
		length int
	}

	node struct {
		value interface{}
		prev  *node
	}
)

// Create a new Stack
func NewStack() *Stack {
	return &Stack{nil, 0}
}

// Len return the number of items in the Stack
func (s *Stack) Len() int {
	return s.length
}

// Peek views the top item on the Stack
func (s *Stack) Peek() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.top.value
}

// Pop the top item of the Stack and return it
func (s *Stack) Pop() interface{} {
	if s.length == 0 {
		return nil
	}

	n := s.top
	s.top = n.prev
	s.length--
	return n.value
}

// Push a value onto the top of the Stack
func (s *Stack) Push(value interface{}) {
	n := &node{value, s.top}
	s.top = n
	s.length++
}
