package lettercasepermutation

import (
	"reflect"
	"testing"
)

func TestPermutations(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected []string
	}{
		`first`: {
			input: "a1b2",
			expected: []string{
				"a1b2",
				"a1B2",
				"A1b2",
				"A1B2",
			},
		},
	}

	for name, tt := range tests {
		inp := tt.input
		exp := tt.expected
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := letterCasePermutation(inp)
			if !reflect.DeepEqual(got, exp) {
				t.Errorf("got: %v, expected: %v", got, exp)
			}
		})
	}
}
