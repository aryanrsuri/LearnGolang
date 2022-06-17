package main

import (
	"fmt"
)

func main() {
	const s string = "Hello World"

	fmt.Println(" Length: ", len(s))

	for _, rune := range s {
		// fmt.Printf(" %#U is at index %d: \n", rune, i)
		fmt.Println(string(rune))
	}

	const s1 string = "$dljsfd"
	var runearr = []rune(s1)
	fmt.Println(runearr)
}
