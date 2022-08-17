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

	fmt.Println()

	var a = 5.67
	var c = &a
	//  c la value a
	// c la &a = id(a)
	//  c la id(p) itself

	fmt.Println("Value stored in variable a = ", a)
	fmt.Println("Address of variable a = ", &a)
	fmt.Println("c saved id of a: ",  c)
	fmt.Println("value *c: ", * c)

	fmt.Println()
	// You can create a pointer using the built-in new() function
	ptr := new(int) // Pointer to an int
	*ptr = 100

	fmt.Printf("Ptr = %#x, *Ptr value = %d\n", ptr, *ptr)

	fmt.Println()

	var d = 75
	var p1 = &d
	var p2 = &d

	if p1 == p2 {
		fmt.Println("Both pointers p1 and p2 point to the same variable.")
	}
}
