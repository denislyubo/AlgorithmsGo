package longestsubstringwithoutrepeatingchars

import (
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := map[string]struct {
		input    string
		expected int
	}{
		`first`: {
			input:    "abcabcbb",
			expected: 3,
		},
		`second`: {
			input:    "h",
			expected: 1,
		},
		`third`: {
			input:    "aab",
			expected: 2,
		},
		`forth`: {
			input:    "pwwkew",
			expected: 3,
		},
	}

	for name, tt := range tests {
		got := lengthOfLongestSubstring(tt.input)
		exp := tt.expected
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got != exp {
				t.Errorf("got: %d, expected: %d", got, exp)
			}
		})
	}
}
