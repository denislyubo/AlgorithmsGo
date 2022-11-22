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
			reverseString(tt.input)
			if !reflect.DeepEqual(tt.input, tt.expected) {
				t.Errorf("got: %v, expected: %v", tt.input, tt.expected)
			}
		})
	}
}
