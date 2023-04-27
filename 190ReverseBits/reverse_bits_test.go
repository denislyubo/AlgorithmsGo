package reversebits

import "testing"

func TestReverseBits(t *testing.T) {
	tests := map[string]struct {
		n        uint32
		expected uint32
	}{
		`first`: {
			n:        0b00000010100101000001111010011100,
			expected: 964176192,
		},
		`second`: {
			n:        0b11111111111111111111111111111101,
			expected: 3221225471,
		},
	}

	for name, tt := range tests {
		n := tt.n
		exp := tt.expected
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := reverseBits(n)
			if got != tt.expected {
				t.Errorf("got: %d, expected: %d", got, exp)
			}
		})
	}
}
