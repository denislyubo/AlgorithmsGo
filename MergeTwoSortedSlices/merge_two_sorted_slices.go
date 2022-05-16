package main

import "fmt"

func mergeSorted(a []int, b []int) []int {
	var i, j int

	c := make([]int, len(a)+len(b))

	for k := 0; k < len(c); k++ {
		if i == len(a) {
			c[k] = b[j]
			j++
			continue
		} else if j == len(b) {
			c[k] = a[i]
			i++
			continue
		} else if a[i] < b[j] {
			c[k] = a[i]
			i++
		} else {
			c[k] = b[j]
			j++
		}
	}

	return c
}

func main() {
	r := mergeSorted([]int{1, 7}, []int{2, 9, 12})

	fmt.Printf("%+v", r)
}
