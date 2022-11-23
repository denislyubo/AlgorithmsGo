package main

import "testing"

func TestTwoNumbersSumEquals(t *testing.T) {
	i, j, _ := twoNumbersSumEquals([]int{3, 4, 2, 56, 12, 9, 3, 4}, 68)

	if (i != 56 && j != 12) && (i != 12 && j != 56) {
		t.Error("error")
	}

}
