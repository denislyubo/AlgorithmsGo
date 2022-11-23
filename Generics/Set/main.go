package main

import "fmt"

// Set generic set type
type Set[T comparable] struct {
	m map[T]struct{}
}

// NewSet is constructor
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		m: make(map[T]struct{}),
	}
}

// Insert elem to set
func (s *Set[T]) Insert(elem T) {
	s.m[elem] = struct{}{}
}

// Delete is deletes elem from set
func (s *Set[T]) Delete(elem T) {
	delete(s.m, elem)
}

func main() {
	s := NewSet[int]()

	s.Insert(12)

	s.Insert(15)

	s.Delete(12)

	s.Insert(37)

	fmt.Printf("%+v", s)
}
