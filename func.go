package main

import (
	"fmt"
	// str "strings" // alias
	"math"
	// "errors"
	"phong/tricks"
	// "golang.org/x/exp/constraints"
)

/*	
	// variadic arguments is 1 slice 
	// variadic arguments: phải được khai báo nằm ở sau cùng 
	func hello(a int, args ...int) {  
		// syntax ok
	}

	func hello(args ...int, b int) {  
		// syntax error
	}

*/



/*
	anonymous function; type interface{} == type any
	lambda := func(slice ...interface{}) []any {
		return slice
	}

	#python
	def action(*args):
		return args

	// anonymous func = lambda: dùng lòng khai báo trong các func
	đoạn code: func(name any) any


*/


// variadic arguments == slice
// ...string == []string
// syntax call func có arguments là ...variadic: nmyFunc(slice...) ; 3dot...
func variadic(slice ...string) []string {
		return slice
}

func changeSlice(slice ...string) {
	slice[0] = "GOGO"
}

type Number interface {
 	int | int8 | int16 | int32 | int64 | float64
}

func genericFunc[OK tricks.SuperNumber](x, y OK) OK {
	return x + y
}





func main() {
	tricks.Format("variadic argument is slice")
	fmt.Println(variadic("one", "two", "three"))

	tricks.Format("change variadic==slice")
	slice := []string{"hellow", "world"}
	fmt.Println("before changeSlice: ", slice)
	changeSlice(slice...)
	fmt.Println("after changeSlice(slice...): ", slice)


	tricks.Format("declared anonymous func = lambda func")
	lambda := func(x, y float64) float64 {
		return math.Pow(x, y)
	}
	fmt.Println("lambda: ", lambda(3,5))


	fmt.Println(math.Pow(3.5, 10))

	value := func(){
      fmt.Println("Welcome! to GeeksforGeeks")
  	}
  	value()


  	fmt.Println(genericFunc(1.2,15.0))
  	fmt.Println()
	
}
