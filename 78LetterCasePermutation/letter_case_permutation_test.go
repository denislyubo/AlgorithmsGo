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
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := letterCasePermutation(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got: %v, expected: %v", got, tt.expected)
			}
		})
	}
}
