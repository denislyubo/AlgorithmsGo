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
		t.Run(name, func(t *testing.T) {
			got := lengthOfLongestSubstring(tt.input)
			if got != tt.expected {
				t.Errorf("got: %d, expected: %d", got, tt.expected)
			}
		})
	}
}
