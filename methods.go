package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.height + 2*r.width
}

func main() {
	var r = rect{width: 10, height: 5}
	fmt.Println(r.area(), r.perim())
	var rp = &r
	fmt.Println(rp.area(), rp.perim())
}
