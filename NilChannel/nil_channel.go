package main

import "fmt"

func main() {
	c1 := make(chan int, 5)
	c2 := make(chan int, 5)

	for i := 1; i <= 3; i++ {
		c1 <- i
		c2 <- i
	}

	close(c1)
	close(c2)

	for i := 1; i <= 10; i++ {
		select {
		case x, ok := <-c1:
			if !ok {
				c1 = nil

			} else {
				fmt.Println("c1: ", x)
			}
		case x, ok := <-c2:
			if !ok {
				c2 = nil
			} else {
				fmt.Println("c2: ", x)
			}

		default:
			fmt.Println("NO DATA!")
		}
	}
}
