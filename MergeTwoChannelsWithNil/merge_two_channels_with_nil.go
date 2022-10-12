package main

import (
	"fmt"
)

func merge1(ch1, ch2 <-chan int) <-chan int {
	ch := make(chan int, 1)

	go func() {
		for v := range ch1 {
			ch <- v
		}
		for v := range ch2 {
			ch <- v
		}
		close(ch)
	}()

	return ch
}

func merge2(ch1, ch2 <-chan int) <-chan int {
	ch := make(chan int, 1)

	go func() {
		for {
			select {
			case v := <-ch1:
				ch <- v
			case v := <-ch2:
				ch <- v
			}
		}
		close(ch)
	}()

	return ch
}

func merge3(ch1, ch2 <-chan int) <-chan int {
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

func merge4(ch1, ch2 <-chan int) <-chan int {
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

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		for _, val := range []int{1, 2, 3} {
			ch1 <- val
		}
		close(ch1)
	}()
	go func() {
		for _, val := range []int{4, 5, 6} {
			ch2 <- val
		}
		close(ch2)
	}()

	//for passing static checks
	var _, _, _, _ = merge1, merge2, merge3, merge4

	var mergedCh = merge3(ch1, ch2)

	for val := range mergedCh {
		fmt.Println(val)
	}

	fmt.Println("End")
}
