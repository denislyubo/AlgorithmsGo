package stack

type stack struct {
	data []interface{}
}

func (s *stack) Empty() bool {
	return len(s.data) == 0
}

func (s *stack) Pop() {
	s.data = s.data[:len(s.data)-1]
}

func (s *stack) Push(item interface{}) {
	s.data = append(s.data, item)
}

func (s *stack) Top() interface{} {
	return s.data[len(s.data)-1]
}
