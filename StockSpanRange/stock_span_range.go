package stockspanrange

import "github.com/golang-collections/collections/stack"

func StockSpanRangeBrute(arr []int) []int {
	n := len(arr)
	SR := make([]int, n)
	SR[0] = 1
	for i := 1; i < len(arr); i++ {
		SR[i] = 1
		for j := i - 1; (j >= 0) && (arr[i] >= arr[j]); j-- {
			SR[i]++
		}
	}
	return SR
}

func StockSpanRange(arr []int) []int {
	n := len(arr)
	res := make([]int, n)
	s := stack.New()
	s.Push(0)
	res[0] = 1
	for i := 1; i < n; i++ {
		for s.Len() != 0 && arr[s.Peek().(int)] <= arr[i] {
			s.Pop()
		}
		if s.Len() == 0 {
			res[i] = i + 1
		} else {
			res[i] = i - s.Peek().(int)
		}

		s.Push(i)
	}

	return res
}
