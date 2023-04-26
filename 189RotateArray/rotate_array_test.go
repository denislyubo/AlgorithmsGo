package rotatearray

import (
	"reflect"
	"testing"
)

func TestReverseArray(t *testing.T) {
	tests := map[string]struct {
		inputArray []int
		steps      int
		expected   []int
	}{
		`first`: {
			inputArray: []int{1, 2, 3, 4, 5, 6, 7},
			expected:   []int{7, 6, 5, 4, 3, 2, 1},
		},
		`second`: {
			inputArray: []int{-1},
			steps:      2,
			expected:   []int{-1},
		},
		`third`: {
			inputArray: []int{-1, -2},
			steps:      3,
			expected:   []int{-2, -1},
		},
		/*`fourth`: {
			inputArray: []int{-1},
			steps:      3,
			expected:   []int{1},
		},*/
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			reverse(tt.inputArray)
			if !reflect.DeepEqual(tt.inputArray, tt.expected) {
				t.Errorf("got: %v, expected: %v", tt.inputArray, tt.expected)
			}
		})
	}
}

func TestRotateArray(t *testing.T) {
	tests := map[string]struct {
		inputArray []int
		steps      int
		expected   []int
	}{
		`first`: {
			inputArray: []int{1, 2, 3, 4, 5, 6, 7},
			steps:      3,
			expected:   []int{5, 6, 7, 1, 2, 3, 4},
		},
		`second`: {
			inputArray: []int{-1, -100, 3, 99},
			steps:      2,
			expected:   []int{3, 99, -1, -100},
		},
		`third`: {
			inputArray: []int{-1},
			steps:      2,
			expected:   []int{-1},
		},
		`fourth`: {
			inputArray: []int{1, 2},
			steps:      2,
			expected:   []int{1, 2},
		},
		`fifth`: {
			inputArray: []int{1, 2, 3},
			steps:      4,
			expected:   []int{3, 1, 2},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			rotate(tt.inputArray, tt.steps)
			if !reflect.DeepEqual(tt.inputArray, tt.expected) {
				t.Errorf("got: %v, expected: %v", tt.inputArray, tt.expected)
			}
		})
	}
}
