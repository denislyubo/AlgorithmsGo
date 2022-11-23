package main

import "fmt"

type byteSize interface {
	byte | int8
}

func printByte[T byteSize](i T) {
	fmt.Println(i)
}

func main() {
	printByte(uint8(6))
	//printByte(6)
}
