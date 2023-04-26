package triangle

import "testing"

func TestTriangle(t *testing.T) {
	tests := map[string]struct {
		input    [][]int
		expected int
	}{
		`first`: {
			input: [][]int{
				{2},
				{3, 4},
				{6, 5, 7},
				{4, 1, 8, 3},
			},
			expected: 11,
		},
		`second`: {
			input: [][]int{
				{-10},
			},
			expected: -10,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := minimumTotal(tt.input)
			exp := tt.expected
			if got != tt.expected {
				t.Errorf("got: %d, expected: %d", got, exp)
			}
		})
	}
}
