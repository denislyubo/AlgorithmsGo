package FindCelebrity

import (
	"reflect"
	"testing"
)

func TestFindCelebrity(t *testing.T) {
	tests := map[string]struct {
		input    [][]bool
		expected int
	}{
		//`first`: {
		//	input: [][]bool{
		//		{true, true, true},
		//		{false, true, false},
		//		{true, true, true},
		//	},
		//	expected: 1,
		//},
		`second`: {
			input: [][]bool{
				{true, true, true},
				{false, true, false},
				{true, false, true},
			},
			expected: -1,
		},
	}

	for name, tt := range tests {
		inp := tt.input
		exp := tt.expected
		t.Run(name, func(t *testing.T) {
			got := findCelebrity(inp)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got: %v, expected: %v", got, exp)
			}
		})
	}
}
