package climbingstairs

import (
	"testing"
)

func TestClimbingStairs(t *testing.T) {
	tests := map[string]struct {
		n        int
		expected int
	}{
		`first`: {
			n:        2,
			expected: 2,
		},
		`second`: {
			n:        3,
			expected: 3,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := climbStairs(tt.n)
			if got != tt.expected {
				t.Errorf("got: %d, expected: %d", got, tt.expected)
			}
		})
	}
}
