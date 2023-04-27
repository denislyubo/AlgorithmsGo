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
		grid := tt.grid
		exp := tt.expected
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := maxAreaOfIsland(grid)
			if !reflect.DeepEqual(exp, got) {
				t.Errorf("got: %v, expected: %v", got, exp)
			}
		})
	}
}
