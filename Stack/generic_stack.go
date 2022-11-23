package stack

// GenStack is data struct for stack
type GenStack[T any] struct {
	data []T
}

// Empty returns if stack empty
func (s *GenStack[T]) Empty() bool {
	return len(s.data) == 0
}

// Pop is removing top element from stack
func (s *GenStack[T]) Pop() {
	s.data = s.data[:len(s.data)-1]
}

// Push push item on top of stack
func (s *GenStack[T]) Push(item T) {
	s.data = append(s.data, item)
}

// Top returns top element, but not removing one from stack
func (s *GenStack[T]) Top() T {
	return s.data[len(s.data)-1]
}
