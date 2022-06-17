package main

import (
	"fmt"
	"time"
)

// lightweight htread of execution
// the two go functions are running asynchronously in sep go routines
func f(test string) {
	for i := 0; i < 3; i++ {
		fmt.Println(test, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
