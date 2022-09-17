package main

import (
	"fmt"
	// "strings"
	"math"


	"golang.org/x/exp/constraints"

	"phong/tricks"
)
/*
	NOT update , thuộc nhóm này có thể la immutable 
	arr := [5]int{1,2,3,4,5}  == tuple(1,2,3,4,5)

	update OK , các types thuộc nhóm default value được gán sau = nil, rất có thể la mutable 
	slice := []int{1,2,3,4,5}  == list(1,2,3,4,5)
	Map := make(map[string]string)  == dict(name="thong", city="long an")


	var varName int float64 string struct , ... khi declared không gán value, thì value default = 0 ""
	empty của type đó

	type Human struct { 
		name string
		age int
	}

	var (
		myint int
		myfloat64 float64
		mystring string
	)
	fmt.Println(myint)				// 0
	fmt.Println(myfloat64)		// 0.0
	fmt.Println(mystring)			// ""
	fmt.Println(human)			// {"": 0}


	các types default = nil khi không gán value: map channel interface{} slices pointer function
	var varNameP *int *string, *struct ..., nếu không gán &varNameNaoDo đó, thì default = nil

	var (
		myint *int
		myfloat64 *float64
		mystring *string
		myStruct *Human
	)

	fmt.Println(myint)				// nil
	fmt.Println(myfloat64)		// nil
	fmt.Println(mystring)			// nil
	fmt.Println(human)			// nil


	result = new(int)  // result is pointer of &int

	var temp int   // declare an int type variable
	var result *int //  declare a pointer to int
	result = &temp 


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
