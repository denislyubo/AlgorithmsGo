package stockmaxprofit

import "testing"

func TestStockMaxProfit(t *testing.T) {
	tests := map[string]struct {
		inputArray []int
		expected   int
	}{
		`first`: {
			inputArray: []int{7, 4, 3},
			expected:   0,
		},
		`second`: {
			inputArray: []int{7, 4, 3, 7, 9, 1},
			expected:   6,
		},
	}

	for name, tt := range tests {
		var inp = append([]int(nil), tt.inputArray...)
		exp := tt.expected
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			res := maxProfit(inp)
			if res != exp {
				t.Errorf("got: %v, expected: %v", res, exp)
			}
		})
	}
}
