package stack

type GenStack[T any] struct {
	data []T
}

func (s *GenStack[T]) Empty() bool {
	return len(s.data) == 0
}

func (s *GenStack[T]) Pop() {
	s.data = s.data[:len(s.data)-1]
}

func (s *GenStack[T]) Push(item T) {
	s.data = append(s.data, item)
}

func (s *GenStack[T]) Top() T {
	return s.data[len(s.data)-1]
}
