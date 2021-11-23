package stack

import "testing"

func TestEmptyStack(t *testing.T) {
	s := &stack{}

	s.Push(1)

	if s.Empty() {
		t.Error("stack isn't empty")
	}

	s.Pop()

	if !s.Empty() {
		t.Error("stack should be empty")
	}
}
