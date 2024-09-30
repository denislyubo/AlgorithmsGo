package _2MaxProfit

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{1, 2, 3},
			expected: 2,
		},
		{
			input:    []int{3, 2, 4},
			expected: 2,
		},
		{
			input:    []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},

		{
			input:    []int{7, 6, 4, 3, 1},
			expected: 0,
		},
		{
			input:    []int{2, 7, 6, 4, 3, 19},
			expected: 17,
		},
	}

	for num, tt := range tests {
		var inp = append([]int(nil), tt.input...)
		exp := tt.expected
		t.Run(fmt.Sprintf("Test#%d", num), func(t *testing.T) {
			t.Parallel()
			got := maxProfit(inp)
			if !reflect.DeepEqual(got, exp) {
				t.Errorf("got: %v, expected: %v", got, exp)
			}
		})
	}
}
