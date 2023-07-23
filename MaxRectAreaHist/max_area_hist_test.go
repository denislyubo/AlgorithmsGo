package maxrectareahist

import (
	"testing"
)

func TestGetMaxArea(t *testing.T) {
	tests := map[string]struct {
		inputArray []int
		expected   int
	}{
		`first`: {
			inputArray: []int{7, 4, 3},
			expected:   9,
		},
		`second`: {
			inputArray: []int{7, 4, 3, 7, 9, 1},
			expected:   15,
		},
	}

	for name, tt := range tests {
		var inp = append([]int(nil), tt.inputArray...)
		exp := tt.expected
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			res := GetMaxArea2(inp)
			if res != exp {
				t.Errorf("got: %v, expected: %v", res, exp)
			}
		})
	}
}
