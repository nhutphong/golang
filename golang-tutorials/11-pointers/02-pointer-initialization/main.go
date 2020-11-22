package main

import "fmt"

func main() {
	var a = 5.67
	var p = &a
	// *p la value a
	// p la &a = id(a)
	// &p la id(p) itself

	fmt.Println("Value stored in variable a = ", a)
	fmt.Println("Address of variable a = ", &a)
	fmt.Println("id p: ", p)
	fmt.Println("value *p: ", *p)
}
