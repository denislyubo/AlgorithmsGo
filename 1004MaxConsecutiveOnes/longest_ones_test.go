package maxconsecutiveones

import "testing"

func TestLongestOnes(t *testing.T) {

	tests := map[string]struct {
		nums     []int
		k        int
		expected int
	}{
		`first`: {
			nums:     []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
			k:        2,
			expected: 6,
		},
		`second`: {
			nums:     []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
			k:        3,
			expected: 10,
		},
	}

	for name, tt := range tests {
		nums := append([]int(nil), tt.nums...)
		k := tt.k
		exp := tt.expected
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := longestOnes(nums, k)
			if got != exp {
				t.Errorf("got: %d, expected: %d", got, exp)
			}
		})
	}

}
