package reverseworldsinstring3

import (
	"reflect"
	"testing"
)

func TestReverseWords(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected string
	}{
		`first`: {
			input:    "Let's take LeetCode contest",
			expected: "s'teL ekat edoCteeL tsetnoc",
		},
	}

	for name, tt := range tests {
		inp := tt.input
		exp := tt.expected
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := reverseWords(inp)
			if !reflect.DeepEqual(got, exp) {
				t.Errorf("got: %v, expected: %v", got, exp)
			}
		})
	}
}
