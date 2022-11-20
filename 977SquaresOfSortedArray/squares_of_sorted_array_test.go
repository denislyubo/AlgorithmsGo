package squaresofsortedarray

import (
	"reflect"
	"testing"
)

func TestSquareOfSortedArray(t *testing.T) {
	tests := map[string]struct {
		input    []int
		expected []int
	}{
		`first`: {
			input:    []int{-4, -1, 0, 3, 10},
			expected: []int{0, 1, 9, 16, 100},
		},
		`second`: {
			input:    []int{-7, -3, 2, 3, 11},
			expected: []int{4, 9, 9, 49, 121},
		},
		`third`: {
			input:    []int{-1, -2},
			expected: []int{1, 4},
		},
		`fourth`: {
			input:    []int{-1},
			expected: []int{1},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := sortedSquares(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got: %v, expected: %v", got, tt.expected)
			}
		})
	}
}
