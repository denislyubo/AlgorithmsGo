package main

import (
	"fmt"
)

func merge0(ch1, ch2 <-chan int) <-chan int {
	ch := make(chan int, 1)

	go func() {
		for ch1 != nil || ch2 != nil {
			select {
			case v, open := <-ch1:
				if !open {
					ch1 = nil
					break
				}
				ch <- v
			case v, open := <-ch2:
				if !open {
					ch2 = nil
					break
				}
				ch <- v
			}
		}
		close(ch)
	}()

	return ch
}

func merge(ch1, ch2 <-chan int) <-chan int {
	ch := make(chan int, 1)
	ch1Closed := false
	ch2Closed := false

	go func() {
		for {
			select {
			case v, open := <-ch1:
				if !open {
					ch1Closed = true
					break
				}
				ch <- v
			case v, open := <-ch2:
				if !open {
					ch2Closed = true
					break
				}
				ch <- v
			}
			if ch1Closed && ch2Closed {
				close(ch)
				return
			}
		}

	}()

	return ch
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for _, v := range []int{1, 2, 3} {
			ch1 <- v
		}
	}()

	go func() {
		for _, v := range []int{4, 5, 6, 7, 8, 9} {
			ch2 <- v
		}
	}()

	var mergedCh = merge(ch1, ch2)

	for val := range mergedCh {
		fmt.Println(val)
	}

	fmt.Println("End")
}
