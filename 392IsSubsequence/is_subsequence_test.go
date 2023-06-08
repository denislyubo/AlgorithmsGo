package issubsequence

import "testing"

func TestIsSubsequence(t *testing.T) {
	tests := map[string]struct {
		s1       string
		s2       string
		expected bool
	}{
		`first`: {
			s1:       "ab",
			s2:       "ahgghghb",
			expected: true,
		},
		`second`: {
			s1:       "arda",
			s2:       "eidbdaarooo",
			expected: false,
		},
		`third`: {
			s1:       "a",
			s2:       "ab",
			expected: true,
		},
		`fourth`: {
			s1:       "adc",
			s2:       "dcda",
			expected: false,
		},
		`fifth`: {
			s1:       "",
			s2:       "dcda",
			expected: false,
		},
	}

	for name, tt := range tests {
		s1 := tt.s1
		s2 := tt.s2
		exp := tt.expected
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := isSubsequence(s1, s2)
			if got != exp {
				t.Errorf("got: %t, expected: %t", got, exp)
			}
		})
	}

}
