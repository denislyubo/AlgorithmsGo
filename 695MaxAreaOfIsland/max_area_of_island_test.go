package maxareaofisland

import (
	"reflect"
	"testing"
)

func TestFloodFill(t *testing.T) {
	tests := map[string]struct {
		grid     [][]int
		expected int
	}{
		`first`: {
			grid: [][]int{
				{1, 1, 1},
				{1, 1, 0},
				{1, 0, 1},
			},
			expected: 6,
		},
		`second`: {
			grid: [][]int{
				{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
				{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
			},
			expected: 6,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := maxAreaOfIsland(tt.grid)
			exp := tt.expected
			if !reflect.DeepEqual(tt.expected, got) {
				t.Errorf("got: %v, expected: %v", got, exp)
			}
		})
	}
}
