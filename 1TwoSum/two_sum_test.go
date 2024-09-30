package _TwoSum

import (
	"reflect"
	"sort"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := map[string]struct {
		input    []int
		target   int
		expected []int
	}{
		`first`: {
			input:    []int{1, 2, 3},
			target:   3,
			expected: []int{0, 1},
		},
		`second`: {
			input:    []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
	}

	for name, tt := range tests {
		var inp = append([]int(nil), tt.input...)
		exp := tt.expected
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := twoSum(inp, tt.target)
			if !reflect.DeepEqual(got, exp) {
				t.Errorf("got: %v, expected: %v", got, exp)
			}
		})
	}
}

func TestTwoSum1(t *testing.T) {
	tests := map[string]struct {
		input    []int
		target   int
		expected []int
	}{
		`first`: {
			input:    []int{1, 2, 3},
			target:   3,
			expected: []int{0, 1},
		},
		`second`: {
			input:    []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},

		`third`: {
			input:    []int{4, 2, 2},
			target:   4,
			expected: []int{1, 2},
		},
	}

	for name, tt := range tests {
		var inp = append([]int(nil), tt.input...)
		exp := append([]int(nil), tt.expected...)
		target := tt.target
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := twoSum1(inp, target)
			sort.Ints(got)
			if !reflect.DeepEqual(got, exp) {
				t.Errorf("got: %v, expected: %v", got, exp)
			}
		})
	}
}

func TestTwoSum2(t *testing.T) {
	tests := map[string]struct {
		input    []int
		target   int
		expected []int
	}{
		`first`: {
			input:    []int{1, 2, 3},
			target:   3,
			expected: []int{0, 1},
		},
		`second`: {
			input:    []int{2, 3, 4},
			target:   6,
			expected: []int{0, 2},
		},

		`third`: {
			input:    []int{2, 2, 4},
			target:   4,
			expected: []int{0, 1},
		},
	}

	for name, tt := range tests {
		var inp = append([]int(nil), tt.input...)
		exp := append([]int(nil), tt.expected...)
		target := tt.target
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := twoSum2(inp, target)
			sort.Ints(got)
			if !reflect.DeepEqual(got, exp) {
				t.Errorf("got: %v, expected: %v", got, exp)
			}
		})
	}
}

func TestTwoSum3(t *testing.T) {
	tests := map[string]struct {
		input    []int
		target   int
		expected []int
	}{
		`first`: {
			input:    []int{1, 2, 3},
			target:   3,
			expected: []int{0, 1},
		},
		`second`: {
			input:    []int{2, 3, 4},
			target:   6,
			expected: []int{0, 2},
		},

		`third`: {
			input:    []int{2, 2, 4},
			target:   4,
			expected: []int{0, 1},
		},
	}

	for name, tt := range tests {
		var inp = append([]int(nil), tt.input...)
		exp := append([]int(nil), tt.expected...)
		target := tt.target
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := twoSum3(inp, target)
			sort.Ints(got)
			if !reflect.DeepEqual(got, exp) {
				t.Errorf("got: %v, expected: %v", got, exp)
			}
		})
	}
}
