package queue

// Queue is data struct for Queue
type Queue struct {
	data []any
}

// Empty returns if Queue empty
func (s *Queue) Empty() bool {
	return len(s.data) == 0
}

// Size returns queue size
func (s *Queue) Size() int {
	return len(s.data)
}

// PopFront is removing first element from Queue
func (s *Queue) PopFront() {
	s.data = s.data[1:len(s.data)]
}

// PushBack pushes item to the end of Queue
func (s *Queue) PushBack(item any) {
	s.data = append(s.data, item)
}

// Top returns top element, but not removing one from Queue
func (s *Queue) Top() any {
	return s.data[0]
}
