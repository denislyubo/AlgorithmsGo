package matrix

import (
	"reflect"
	"testing"
)

func TestMatrix(t *testing.T) {
	tests := map[string]struct {
		input    [][]int
		expected [][]int
	}{
		`first`: {
			input: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
			expected: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{0, 0, 0},
			},
		},
		`second`: {
			input: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{1, 1, 1},
			},
			expected: [][]int{
				{0, 0, 0},
				{0, 1, 0},
				{1, 2, 1},
			},
		},
	}

	for name, tt := range tests {
		var inp = make([][]int, 3)
		for j := 0; j < len(tt.input); j++ {
			inp[j] = append([]int(nil), tt.input[j]...)
		}
		exp := tt.expected
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := updateMatrix(inp)
			if !reflect.DeepEqual(got, exp) {
				t.Errorf("got: %v, expected: %v", inp, exp)
			}
		})
	}
}
