package missingnumber

import "testing"

func TestMissingNumber(t *testing.T) {
	if missingNumber([]int{3, 0, 1}) != 2 {
		t.Error(`missingNumber([]int{3, 0, 1}) != 2`)
	}

	if missingNumber([]int{0, 1}) != 2 {
		t.Error(`missingNumber([]int{0, 1}) != 2`)
	}

	if missingNumber([]int{1}) != 0 {
		t.Error(`missingNumber([]int{1}) != 0`)
	}
}
