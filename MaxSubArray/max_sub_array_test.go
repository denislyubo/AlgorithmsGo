package maxsubarray

import (
	"reflect"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	tests := map[string]struct {
		inputArray []int
		expected   []int
	}{
		`first`: {
			inputArray: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			expected:   []int{4, -1, 2, 1},
		},
		`second`: {
			inputArray: []int{-2, -2, -1, -1, -1, -1, -1, -5, -9},
			expected:   []int{-1},
		},
	}

	for name, tt := range tests {
		var inp = append([]int(nil), tt.inputArray...)
		exp := append([]int(nil), tt.expected...)
		t.Run(name, func(t *testing.T) {
			//t.Parallel()
			res := maxSubArray(inp)
			if !reflect.DeepEqual(res, exp) {
				t.Errorf("got: %v, expected: %v", res, exp)
			}
		})
	}
}
