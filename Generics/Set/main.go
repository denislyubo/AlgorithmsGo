package main

import "fmt"

type Set[T comparable] struct {
	m map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{
		m: make(map[T]struct{}),
	}
}

func (s *Set[T]) Insert(elem T) {
	s.m[elem] = struct{}{}
}

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
