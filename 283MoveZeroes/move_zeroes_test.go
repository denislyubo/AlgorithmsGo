package movezeroes

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := map[string]struct {
		inputArray []int
		expected   []int
	}{
		`first`: {
			inputArray: []int{0, 1, 3, 0, 12},
			expected:   []int{1, 3, 12, 0, 0},
		},
		`second`: {
			inputArray: []int{-1},
			expected:   []int{-1},
		},
		`third`: {
			inputArray: []int{0},
			expected:   []int{0},
		},
		/*`fourth`: {
			inputArray: []int{-1},
			expected:   []int{1},
		},*/
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			moveZeroes(tt.inputArray)
			if !reflect.DeepEqual(tt.inputArray, tt.expected) {
				t.Errorf("got: %v, expected: %v", tt.inputArray, tt.expected)
			}
		})
	}
}
