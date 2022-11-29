package combinations

import (
	"reflect"
	"testing"
)

func TestMatrix(t *testing.T) {
	tests := map[string]struct {
		n, k     int
		expected [][]int
	}{
		`first`: {
			n: 4,
			k: 2,
			expected: [][]int{
				{1, 2},
				{1, 3},
				{1, 4},
				{2, 3},
				{2, 4},
				{3, 4},
			},
		},
		`second`: {
			n: 1,
			k: 1,
			expected: [][]int{
				{1},
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := combine(tt.n, tt.k)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("got: %v, expected: %v", got, tt.expected)
			}
		})
	}
}
