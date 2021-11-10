package binarysearch

import "testing"

func TestSearch(t *testing.T) {
	if search([]int{1, 3, 7, 10}, 3) != 1 {
		t.Error(`search([]int{1, 3, 7, 10}, 3) != 1`)
	}

	if search([]int{1, 3, 7, 10}, 11) != -1 {
		t.Error(`search([]int{1, 3, 7, 10}, 3) != -11`)
	}

	if search([]int{1}, 1) != 0 {
		t.Error(`search([]int{1}, 1) != 0`)
	}

	if search([]int{1,7}, 7) != 1 {
		t.Error(`search([]int{1,7}, 7) != 1`)
	}

	if search([]int{}, 7) != -1 {
		t.Error(`earch([]int{}, 7) != -1`)
	}
}
