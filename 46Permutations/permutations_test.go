package permutations

import (
	"reflect"
	"testing"
)

func TestPermutations(t *testing.T) {
	tests := map[string]struct {
		input    []int
		expected [][]int
	}{
		`first`: {
			input: []int{1, 2, 3},
			expected: [][]int{
				{1, 2, 3},
				{1, 3, 2},
				{2, 1, 3},
				{2, 3, 1},
				{3, 2, 1},
				{3, 1, 2},
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := permute(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got: %v, expected: %v", got, tt.expected)
			}
		})
	}
}
