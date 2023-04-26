package rottingoranges

import (
	"testing"
)

func TestRottingOranges(t *testing.T) {
	tests := map[string]struct {
		input    [][]int
		expected int
	}{
		`first`: {
			input: [][]int{
				{2, 1, 1},
				{1, 1, 0},
				{0, 1, 1},
			},
			expected: 4,
		},
		`second`: {
			input: [][]int{
				{2, 1, 1},
				{0, 1, 1},
				{1, 0, 1},
			},
			expected: -1,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			inp := tt.input
			got := orangesRotting(inp)
			exp := tt.expected
			if got != exp {
				t.Errorf("got: %v, expected: %v", got, exp)
			}
		})
	}
}
