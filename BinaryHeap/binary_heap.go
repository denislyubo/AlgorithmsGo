package binaryheap

// BinaryHeap - storage struct
type BinaryHeap struct {
	data []int
}

// NewBinaryHeap is constructor
func NewBinaryHeap() *BinaryHeap {
	return &BinaryHeap{}
}

// Size returns heap size
func (bh *BinaryHeap) Size() int {
	return len(bh.data)
}

// Add is adding element to heap
func (bh *BinaryHeap) Add(val int) {
	(*bh).data = append(bh.data, val)

	i := bh.Size() - 1
	parent := (i - 1) / 2

	for i > 0 && bh.data[parent] < bh.data[i] {
		bh.data[i], bh.data[parent] = bh.data[parent], bh.data[i]

		i = parent
		parent = (i - 1) / 2
	}
}

func (bh *BinaryHeap) heapify(i int) {
	var leftChild, rightChild, largestChild int

	for {
		leftChild = 2*i + 1
		rightChild = 2*i + 2
		largestChild = i

		if leftChild < bh.Size() && bh.data[leftChild] > bh.data[largestChild] {
			largestChild = leftChild
		}

		if rightChild < bh.Size() && bh.data[rightChild] > bh.data[largestChild] {
			largestChild = rightChild
		}

		if largestChild == i {
			break
		}

		bh.data[i], bh.data[largestChild] = bh.data[largestChild], bh.data[i]
	}
}

// BuildHeap is function for heap ordering of underlying array
func (bh *BinaryHeap) BuildHeap(sourceArray []int) {
	(*bh).data = sourceArray
	for i := bh.Size() / 2; i >= 0; i-- {
		bh.heapify(i)
	}
}

// GetMax returns element on top of heap and removes it from heap
func (bh *BinaryHeap) GetMax() int {
	var result int = bh.data[0]
	bh.data[0] = bh.data[bh.Size()-1]
	bh.data = bh.data[0 : bh.Size()-1]
	bh.heapify(0)
	return result
}

// HeapSort is sorting array using heap
func (bh *BinaryHeap) HeapSort(array []int) {
	bh.BuildHeap(array)
	for i := len(array) - 1; i >= 0; i-- {
		array[i] = bh.GetMax()
		bh.heapify(0)
	}
}
