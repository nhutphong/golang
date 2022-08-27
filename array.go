package main

import (
	"fmt"
	"phong/tricks"
)

func main() {
	//  assign( = , := ) cua array la VALUE TYPE = tham tri

	var x [5]int // An array of 5 integers
	fmt.Println(x)

	tricks.Format("slice")
	var slice []int
	fmt.Printf("type is %T", slice)


	tricks.Format("*")
	var y [8]string // An array of 8 strings
	fmt.Println(y)

	var z [3]complex128 // An array of 3 complex numbers
	fmt.Println(z)

	
	// Declaring and initializing an array at the same time
	var a = [5]int{2, 4, 6, 8, 10}
	fmt.Println(a)

	// Short hand declaration for declaring and initializing an array
	b := [5]int{2, 4, 6, 8, 10}
	fmt.Println(b)

	tricks.Format("[...]int{3, 5, 7, 9}  3dot length tuy y")
	// You don't need to initialize all the elements of the array.
	// The un-initialized elements will be assigned the zero value of the corresponding array type
	c := [5]int{2}
	fmt.Println(c)

	// do dai array tuy y
	d := [...]int{3, 5, 7, 9, 11, 13, 17, 19, 22, 50}
	fmt.Println(d)


	// golang la tham tri=value
	a1 := [5]string{"English", "Japanese", "Spanish", "French", "Hindi"}
	a2 := a1 // A copy of the array `a1` is assigned to `a2`

	// a2 thay doi khong lam a1 thay doi
	a2[1] = "German"

	fmt.Println("a1 = ", a1) // The array `a1` remains unchanged
	fmt.Println("a2 = ", a2)


	// Type=struc=class

	// dung make() create array
	// array := make([]int, 10)

	// default for cho array
	// for index := range array

	// for index, value := range array {
		//}
	// for index, value in enumerate(array, 1)  == python

	// ignoze value dung underscore
	// for i, _ := range array {}

	// ignoze index dung underscore
	// for _, value := range array {}

	tricks.Format("for index, value := range array {}")
	daysOfWeek := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	for index, value := range daysOfWeek {
		fmt.Printf("Day %d of week = %s\n", index, value)
	}


	tricks.Format("for _, value := range array {}")
	one := [4]float64{3.5, 7.2, 4.8, 9.5}
	total := float64(0)

	for _, value := range one {
		total = total + value
	}

	fmt.Printf("total of all the elements in array %v = %f\n", a, total)


	tricks.Format("array 2 chieu")
	// 1 parent array chua 2 child array chua 3 item=int
	// child array co the co item it hon len=qui dinh, nhung ko dc > hon
	two := [2][3]int{
		{1, 3, 5},
		{7, 9, 20}, // This comma is necessary
	}
	fmt.Println(two)

	three := [3][4]float64{
		{1, 3},
		{4.5, -3, 7.4, 2},
		{6, 2, 11},
	}
	fmt.Println(three)
}
