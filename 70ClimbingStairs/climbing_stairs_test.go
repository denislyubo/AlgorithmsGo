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
		n := tt.n
		exp := tt.expected
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := climbStairs(n)
			if got != exp {
				t.Errorf("got: %d, expected: %d", got, exp)
			}
		})
	}
}
