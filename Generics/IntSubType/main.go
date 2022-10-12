package main

import "fmt"

type ByteSize interface {
	byte | int8
}

func PrintByte[T ByteSize](i T) {
	fmt.Println(i)
}

func main() {
	PrintByte(uint8(6))
	//PrintByte(6)
}
