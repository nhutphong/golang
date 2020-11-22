package main

import "fmt"

func main() {
	var a = 7.98
	var p = &a
	var pp = &p

	fmt.Println("a = ", a)
	fmt.Println("ID &a = ", &a)

	fmt.Println("p = ", p)
	fmt.Println("ID &p = ", &p)

	fmt.Println("pp = &p = ", pp)

	// Dereferencing a pointer to pointer
	fmt.Println("*pp = ", *pp)
	fmt.Println("**pp = ", **pp)
}
