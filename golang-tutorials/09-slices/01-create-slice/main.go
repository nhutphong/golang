package main

import "fmt"

func main() {

	// array := [7]int{3, 5, 7, 9, 11, 13, 17}
	// array co dinh len [7]
	// slice ko co dinh len []
	var s = []int{3, 5, 7, 9, 11, 13, 17} // Creates an array, and returns a slice reference to the array

	// Short hand declaration
	t := []int{2, 4, 8, 16, 32, 64}

	fmt.Println("s = ", s)
	fmt.Println("t = ", t)
}
