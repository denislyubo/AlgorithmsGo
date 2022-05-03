package main

import "fmt"

func mergeSorted(a []int, b []int) (c []int) {
	var i, j int

	for {
		if i >= len(a) {
			j++
		}
		if j >= len(b) {
			i++
		}

		if i >= len(a) && j >= len(b) {
			break
		}

		if a[i] < b[j] {
			c = append(c, a[i])
			i++
		} else {
			c = append(c, b[j])
			j++
		}
	}

	return c
}

func main() {
	r := mergeSorted([]int{1, 7, 9}, []int{2, 12})

	fmt.Printf("%+v", r)
}
