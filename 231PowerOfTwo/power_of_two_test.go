package poweroftwo

import "testing"

func TestPowerOfTwo(t *testing.T) {
	tests := map[string]struct {
		n        int
		expected bool
	}{
		`first`: {
			n:        1,
			expected: true,
		},
		`second`: {
			n:        16,
			expected: true,
		},
		`third`: {
			n:        3,
			expected: false,
		},
		`forth`: {
			n:        8,
			expected: true,
		},
		`fifth`: {
			n:        6,
			expected: false,
		},
	}

	for name, tt := range tests {
		n := tt.n
		exp := tt.expected
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := isPowerOfTwo(n)
			if got != exp {
				t.Errorf("got: %v, expected: %v", got, exp)
			}
		})
	}
}
