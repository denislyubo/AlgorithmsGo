package singlenumber

import (
	"testing"
)

func TestReverseBits(t *testing.T) {
	tests := map[string]struct {
		nums     []int
		expected int
	}{
		`first`: {
			nums:     []int{2, 2, 1},
			expected: 1,
		},
		`second`: {
			nums:     []int{4, 2, 1, 2, 1},
			expected: 4,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := singleNumber(tt.nums)
			if got != tt.expected {
				t.Errorf("got: %v, expected: %v", got, tt.expected)
			}
		})
	}
}
