package main

import (
	"fmt"
	// "strings"
	"math"


	"golang.org/x/exp/constraints"

	"phong/tricks"
)
/*
	func main() khong the declared them 1 func voi syntax basic:
		func change(){ //statement }

	dung anonymou func declared bao nhieu cung duoc
		change := func(){}
		update := func(){}
		delete := func(){}
	
*/

/*
	var number int
	var ptr *int // gan value cho ptr phai la pointer

	ptr = &number

	set value cho pointer
	*ptr = 22 // gan value cho pointer, khi do number=22

	fmt.Println(ptr)	// 1 so nao do trong RAM
	fmt.Println(*ptr)	// 22
	fmt.Println(number) // 22


	def input_python
	var num int
	var name string

	pass pointer &num &name vao func ben duoi
	fmt.Scan(&num) // nhap vo 10
	fmt.Println(num) //10


	Scanf("%d", &num) 		voi format: %c d e f, ..
	scanln(&name)		:khi detected /n stop


*/

// function definition with a pointer argument
func update(num *int) {
  // dereferencing the pointer
  *num = 30

} 

func display() *string {

  message := "Programiz"

  // returns the address of message
  return &message
}
/*
	/ function call
  result := display() 
  fmt.Println("Welcome to", *result)
*/



// Ordered is a type constraint that matches any ordered type.
// An ordered type is one that supports the <, <=, >, and >= operators.
// type Ordered interface {
// 	type int, int8, int16, int32, int64,
// 		uint, uint8, uint16, uint32, uint64, uintptr,
// 		float32, float64,
// 		string
// } 


type Number interface {
    constraints.Integer | constraints.Float
}

//
func test[T Number](a,b T) float64 {
	return math.Pow(float64(a), float64(b))
}


func main() {
	tricks.Format("*")

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


	tricks.Format("change(&name)")

	change := func(name *string) {
		*name = "WWWWW"
	}

	name := "thong"
	change(&name)
	fmt.Println(name)

	
	
	// fmt.Println(add(2,5))
	fmt.Println(test(2,5))
}
