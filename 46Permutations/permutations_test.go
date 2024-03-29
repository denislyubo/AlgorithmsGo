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
		var inp = append([]int(nil), tt.input...)
		exp := tt.expected
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := permute(inp)
			if !reflect.DeepEqual(got, exp) {
				t.Errorf("got: %v, expected: %v", got, exp)
			}
		})
	}
}
