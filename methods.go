package main

import "fmt"

type Rectangle struct {
	len, breadth int
}

func (r Rectangle) Area() int {
	return r.len * r.breadth
}

func main() {
	r := Rectangle{3, 5}

	area := r.Area()

	fmt.Print("Area of the triangle is: ", area)
}
