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
		t.Run(name, func(t *testing.T) {
			got := updateMatrix(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got: %v, expected: %v", tt.input, tt.expected)
			}
		})
	}
}
