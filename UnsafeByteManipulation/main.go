package main

import (
	"fmt"
)

func main() {
	v := 0x11102
	//d := uintptr(unsafe.Pointer(&v)) + 1
	//fmt.Printf("%x\n", d)
	//*(*byte)(unsafe.Pointer(d)) = 0

	fmt.Println(v)
}
