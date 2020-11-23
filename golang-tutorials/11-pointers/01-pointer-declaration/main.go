package main

import "fmt"

func main() {
	var p *int // nil, not assignment value
	fmt.Println("p = ", p)

	var number int = 123
	pointer := &number

	fmt.Println("number = ", number)
	fmt.Println("pointer = ", pointer)
	fmt.Println("*pointer = ", *pointer)

	fmt.Println()
	*pointer = 555
	fmt.Println("assignment: *pointer = 555")
	fmt.Println("*pointer is ", *pointer)
	fmt.Println("number = ", number)
}
