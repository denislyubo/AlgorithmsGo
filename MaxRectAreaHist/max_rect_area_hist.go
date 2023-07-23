package maxrectareahist

import "github.com/golang-collections/collections/stack"

func GetMaxArea(arr []int) int {
	size := len(arr)
	maxArea := -1
	curArea := 0
	minHeight := 0

	for i := 1; i < size; i++ {
		minHeight = arr[i]
		for j := i - 1; j >= 0; j-- {
			if minHeight > arr[j] {
				minHeight = arr[j]
			}

			curArea = minHeight * (i - j + 1)
			if maxArea < curArea {
				maxArea = curArea
			}
		}
	}

	return maxArea
}

func GetMaxArea2(arr []int) int {
	size := len(arr)
	stk := stack.New()
	maxArea := 0
	Top := 0
	TopArea := 0
	i := 0
	for i < size {
		for (i < size) && (stk.Len() == 0 || arr[stk.Peek().(int)] <= arr[i]) {
			stk.Push(i)
			i++
		}
		for stk.Len() != 0 && (i == size || arr[stk.Peek().(int)] > arr[i]) {
			Top = stk.Peek().(int)
			stk.Pop()
			if stk.Len() == 0 {
				TopArea = arr[Top] * i
			} else {
				TopArea = arr[Top] * (i - stk.Peek().(int) - 1)
			}
			if maxArea < TopArea {
				maxArea = TopArea
			}
		}
	}
	return maxArea
}
