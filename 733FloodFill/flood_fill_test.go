package floodfill

import (
	"reflect"
	"testing"
)

func TestFloodFill(t *testing.T) {
	tests := map[string]struct {
		image    [][]int
		sr, sc   int
		color    int
		expected [][]int
	}{
		`first`: {
			image: [][]int{
				{1, 1, 1},
				{1, 1, 0},
				{1, 0, 1},
			},
			sr:    1,
			sc:    1,
			color: 2,
			expected: [][]int{
				{2, 2, 2},
				{2, 2, 0},
				{2, 0, 1},
			},
		},
	}

	for name, tt := range tests {
		i := tt.image
		sr := tt.sr
		sc := tt.sc
		col := tt.color
		exp := tt.expected
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			got := floodFill(i, sr, sc, col)

			if !reflect.DeepEqual(tt.expected, got) {
				t.Errorf("got: %v, expected: %v", got, exp)
			}
		})
	}
}
