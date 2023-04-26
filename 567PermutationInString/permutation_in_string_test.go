package permutationinstring

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := map[string]struct {
		s1       string
		s2       string
		expected bool
	}{
		`first`: {
			s1:       "ab",
			s2:       "eidbaooo",
			expected: true,
		},
		`second`: {
			s1:       "arda",
			s2:       "eidbdaarooo",
			expected: true,
		},
		`third`: {
			s1:       "a",
			s2:       "ab",
			expected: true,
		},
		`fourth`: {
			s1:       "adc",
			s2:       "dcda",
			expected: true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := checkInclusion(tt.s1, tt.s2)
			exp := tt.expected
			if got != tt.expected {
				t.Errorf("got: %t, expected: %t", got, exp)
			}
		})
	}
}
