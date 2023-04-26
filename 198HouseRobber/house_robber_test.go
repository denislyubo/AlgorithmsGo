package houserobber

import "testing"

func TestRob(t *testing.T) {
	tests := map[string]struct {
		input    []int
		expected int
	}{
		`first`: {
			input:    []int{1, 2, 3, 1},
			expected: 4,
		},
		`second`: {
			input:    []int{2, 7, 9, 3, 1},
			expected: 12,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := rob(tt.input)
			exp := tt.expected
			if got != tt.expected {
				t.Errorf("got: %d, expected: %d", got, exp)
			}
		})
	}
}
