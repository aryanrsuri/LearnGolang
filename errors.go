package main

import (
	"errors"
	"fmt"
)

func test1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("im")
	}
	return arg + 3, nil
}

type argerr struct {
	arg  int
	prob string
}

func (e *argerr) Error() string {
	return fmt.Sprintf("%d -- %s ", e.arg, e.prob)
}

func test2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argerr{arg, "im"}

	}
	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if res, err := test1(i); err != nil {
			fmt.Println("test1 failes", err)

		} else {
			fmt.Println("wirjed", res)
		}
	}

	for _, i := range []int{7, 42} {
		if res, err := test2(i); err != nil {
			fmt.Println("test1 failes", err)

		} else {
			fmt.Println("wirjed", res)
		}
	}

	_, err := test2(42)

	if ae, ok := err.(*argerr); ok {
		fmt.Println(ae.arg, ae.prob)
	}

}
