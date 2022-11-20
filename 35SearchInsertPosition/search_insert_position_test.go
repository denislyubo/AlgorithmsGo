package searchinsertposition

import "testing"

func TestSearchInsertPosition(t *testing.T) {
	tests := map[string]struct {
		nums     []int
		target   int
		expected int
	}{
		`target not in slice  of len 1 expected insert first`: {
			nums:     []int{1, 3, 5, 6},
			target:   0,
			expected: 0,
		},
		`target not in slice 1 expected append third`: {
			nums:     []int{5, 6, 8},
			target:   7,
			expected: 2,
		},
		`target not in slice 2 expected append second`: {
			nums:     []int{1, 3, 5, 6},
			target:   2,
			expected: 1,
		},
		`target not in slice expected append last`: {
			nums:     []int{1, 3, 5, 6},
			target:   7,
			expected: 4,
		},
		`target not in slice expected insert first`: {
			nums:     []int{1, 3, 5, 6},
			target:   0,
			expected: 0,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := searchInsert(tt.nums, tt.target)
			if got != tt.expected {
				t.Errorf("got: %d, expected: %d", got, tt.expected)
			}
		})
	}
}
