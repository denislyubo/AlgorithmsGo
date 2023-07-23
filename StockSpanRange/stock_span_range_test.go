package stockspanrange

import (
	"reflect"
	"testing"
)

func TestStockSpanRange(t *testing.T) {
	tests := map[string]struct {
		inputArray []int
		expected   []int
	}{
		`first`: {
			inputArray: []int{7, 6, 5, 4, 3, 4, 5, 6, 9},
			expected:   []int{1, 1, 1, 1, 1, 3, 5, 7, 9},
		},
	}

	for name, tt := range tests {
		var inp = append([]int(nil), tt.inputArray...)
		exp := append([]int(nil), tt.expected...)
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			res := StockSpanRange(inp)
			if !reflect.DeepEqual(exp, res) {
				t.Errorf("got: %v, expected: %v", res, exp)
			}
		})
	}
}

func TestStockSpanRangeBrute(t *testing.T) {
	tests := map[string]struct {
		inputArray []int
		expected   []int
	}{
		`first`: {
			inputArray: []int{7, 6, 5, 4, 3, 4, 5, 6, 9},
			expected:   []int{1, 1, 1, 1, 1, 3, 5, 7, 9},
		},
	}

	for name, tt := range tests {
		var inp = append([]int(nil), tt.inputArray...)
		exp := append([]int(nil), tt.expected...)
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			res := StockSpanRangeBrute(inp)
			if !reflect.DeepEqual(exp, res) {
				t.Errorf("got: %v, expected: %v", res, exp)
			}
		})
	}
}
