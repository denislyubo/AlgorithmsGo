package reversestring

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	tests := map[string]struct {
		input    []byte
		expected []byte
	}{
		`first`: {
			input:    []byte("dexter"),
			expected: []byte("retxed"),
		},
		`second`: {
			input:    []byte("h"),
			expected: []byte("h"),
		},
		`third`: {
			input:    []byte("ha"),
			expected: []byte("ah"),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			reverseString(tt.input)
			inp := tt.input
			exp := tt.expected
			if !reflect.DeepEqual(inp, exp) {
				t.Errorf("got: %v, expected: %v", inp, exp)
			}
		})
	}
}
