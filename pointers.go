package main

import "fmt"

// Go supports pointers

// This func sets the value to zero as a copy of _val different from the one in the calling function
func zeroval(_val int) {
	_val = 0
}

// This func dereferences the pointer from its memory address to the current value at that address
func zeroptr(_ptr *int) {
	*_ptr = 0
}

func main() {
	var i int = 1
	zeroval(i)
	fmt.Println("zeroval: ", i)
	zeroptr(&i)
	fmt.Println("zeroptr: ", i)
	fmt.Println("pointer: ", &i)
}
