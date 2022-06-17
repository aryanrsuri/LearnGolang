package main

import "fmt"

func sum(a, b, c float64) float64 {
	var sum float64 = a + b + c
	return sum

}

func multiplereturn() (int, int) {
	return 3, 67
}

// Variadic Function
func varfunc(nums ...int) {
	fmt.Println(nums)
	var _total int = 0
	for _, num := range nums {
		_total += num
	}
	fmt.Println(_total)

}

// Anonymous functions and their closures

func intSeq() func() int {
	var i int = 0
	return func() int {
		i++
		return i
	}
}

// Recursive Functions

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(sum(1, 2.4, 3))

	_, c := multiplereturn()
	fmt.Println(c)

	var nums = []int{1, 2, 3, 4}
	varfunc(1, 2, 3)
	varfunc(nums...)

	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())

	fmt.Println(fact(7))

	// recursive closures must be defined explicitley
	var fib func(n int) int

	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-1) + fib(n-2)
	}

	fmt.Println(fib(7))
}
