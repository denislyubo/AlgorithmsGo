package main

import "fmt"

// m = [3, 5, 2, 1, 7, 9, 3, 4]    k== 9
//1. i = 0
//2. 9  - m[i] = 6
//3. i = 1 ... 7

func twoNumbersSumEquals(mas []int, k int) (int, int, error) {
	m := make(map[int]int)

	var dest int = 0

	for i := range mas {
		m[mas[i]]++
	}

	for j := range mas {
		dest = k - mas[j]
		if val, ok := m[dest]; ok {
			if 2*mas[j] == k {
				if val > 1 {
					return mas[j], dest, nil
				} else {
					continue
				}
			}

			return mas[j], dest, nil
		}
	}

	return 0, 0, fmt.Errorf("no such pair")
}

func main() {

	//i, j, err := twoNumbersSumEquals([]int{3, 5, 2, 1, 7, 9, 3, 4}, 9)

	i, j, err := twoNumbersSumEquals([]int{4, 3, 3, 3, 3, 3, 3, 3}, 8)

	fmt.Printf("i=%d j=%d err=%v", i, j, err)
}
