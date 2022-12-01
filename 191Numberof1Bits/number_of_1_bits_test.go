package numberof1its

import "testing"

func TestHammingWeight(t *testing.T) {
	tests := map[string]struct {
		n        uint32
		expected int
	}{
		`first`: {
			n:        1,
			expected: 1,
		},
		`second`: {
			n:        16,
			expected: 1,
		},
		`third`: {
			n:        3,
			expected: 2,
		},
		`forth`: {
			n:        0b00000000000000000000000000001011,
			expected: 3,
		},
		`fifth`: {
			n:        0b11111111111111111111111111111101,
			expected: 31,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := hammingWeight(tt.n)
			if got != tt.expected {
				t.Errorf("got: %d, expected: %d", got, tt.expected)
			}
		})
	}
}
