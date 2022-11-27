package queue

// GenQueue is data struct for Queue
type GenQueue[T any] struct {
	data []T
}

// Empty returns if Queue empty
func (s *GenQueue[T]) Empty() bool {
	return len(s.data) == 0
}

// Size returns queue size
func (s *GenQueue[T]) Size() int {
	return len(s.data)
}

// PopFront is removing first element from Queue
func (s *GenQueue[T]) PopFront() {
	s.data = s.data[1:len(s.data)]
}

// PushBack pushes item to the end of Queue
func (s *GenQueue[T]) PushBack(item T) {
	s.data = append(s.data, item)
}

// Top returns top element, but not removing one from Queue
func (s *GenQueue[T]) Top() T {
	return s.data[0]
}
