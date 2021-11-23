package binaryheap

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBinaryHeap(t *testing.T) {
	heap := NewBinaryHeap()

	raw := []int{17, 3, 7, 2}
	heap.HeapSort(raw)

	if !reflect.DeepEqual(raw, []int{2, 3, 7, 17}) {
		t.Error(fmt.Sprint(raw) + " !=[]int{2, 3, 7, 17}")
	}
}

func TestBinaryHeapAdd(t *testing.T) {
	heap := NewBinaryHeap()

	heap.Add(28)

	if heap.data[0] != 28 {
		t.Error("Add error")
	}

}

func TestBinaryHeapAddTwoNumbers(t *testing.T) {
	heap := NewBinaryHeap()

	heap.Add(28)
	heap.Add(7)

	if !reflect.DeepEqual(heap.data, []int{28, 7}) {
		t.Error(fmt.Sprint(heap.data) + " !=[]int{28, 7}")
	}

}

func TestBinaryHeapAddThreeNumbers(t *testing.T) {
	heap := NewBinaryHeap()

	heap.Add(7)
	heap.Add(9)
	heap.Add(11)

	if !reflect.DeepEqual(heap.data, []int{11, 7, 9}) {
		t.Error(fmt.Sprint(heap.data) + " !=[]int{11, 7, 9}")
	}

}
