package twosum2inputarraysorted

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := map[string]struct {
		numbers  []int
		target   int
		expected []int
	}{
		`first`: {
			numbers:  []int{2, 7, 11, 15},
			target:   9,
			expected: []int{1, 2},
		},
		`second`: {
			numbers:  []int{2, 7, 11, 15, 18, 27},
			target:   26,
			expected: []int{3, 4},
		},
		`third`: {
			numbers:  []int{2, 7},
			target:   15,
			expected: []int{},
		},
		`forth`: {
			numbers:  []int{2},
			target:   2,
			expected: []int{},
		},
		`fifth`: {
			numbers:  []int{},
			target:   2,
			expected: []int{},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := twoSum(tt.numbers, tt.target)
			exp := tt.expected
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got: %v, expected: %v", got, exp)
			}
		})
	}
}
