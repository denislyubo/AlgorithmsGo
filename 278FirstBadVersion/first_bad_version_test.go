package firstbadversion

import "testing"

func TestFirstBadVersion(t *testing.T) {
	getIsBadVersion := func(k int) func(int) bool {
		return func(i int) bool {
			return i >= k
		}
	}

	_ = getIsBadVersion(4)

	tests := map[string]struct {
		n        int
		bad      int
		expected int
	}{
		`n_5 bad_4`: {
			n:        5,
			bad:      4,
			expected: 4,
		},
		`n_50 bad_23`: {
			n:        50,
			bad:      23,
			expected: 23,
		},

		`n_500 bad_499`: {
			n:        500,
			bad:      499,
			expected: 499,
		},
		`n_5000 bad_2`: {
			n:        5000,
			bad:      2,
			expected: 2,
		},
		`n_1 bad_1`: {
			n:        1,
			bad:      1,
			expected: 1,
		},
		`n_2 bad_1`: {
			n:        2,
			bad:      1,
			expected: 1,
		},
		`n_3 bad_1`: {
			n:        3,
			bad:      1,
			expected: 1,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			isBadVersion = getIsBadVersion(tt.bad)
			got := firstBadVersion(tt.n)
			if got != tt.expected {
				t.Errorf("got: %d, expected: %d", got, tt.expected)
			}
		})
	}
}
