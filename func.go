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
	// variadic arguments: phải được khai báo nằm ở sau cùng
	func hello(a int, args ...int) {
		// syntax ok
	}

	func hello(args ...int, b int) {
		// syntax error
	}

*/

// variadic arguments == slice = list
// ...string == []string
// syntax call func có arguments là ...variadic: nmyFunc(slice...) ; 3dot...
func variadic(slice ...string) []string {
	return slice
}

func changeSlice(slice ...string) {
	slice[0] = "GOGO"
}

type Student struct {
	firstName string
	lastName  string
	grade     string
	country   string
}

// named arguments da duoc declared: students, callback; vi the dung =, not declared(:=)
func filter(students []Student, callback func(Student) bool) []Student {
	var students_list []Student
	for _, student := range students {
		if callback(student) == true {
			students_list = append(students_list, student)
		}
	}
	return students_list
}

// simple() return func(a, b int) int
func simple() func(a, b int) int {
	add := func(a, b int) int {
		return a + b
	}
	return add
}

// recursion
func factorial(num int) int {

	// condition to break recursion
	if num == 0 {
		return 1
	} else {

		return num * factorial(num-1)
	}
}

// error vi Pow(float64, float64)
// func test(a, b any) any {
// 	return math.Pow(a,b)
// }

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


	// defined a anonymous function
	  var greet = func() {
	    fmt.Println("Hello, how are you")
	  }

	// function call
	greet()


	// anonymous function return int
	  area := func(length, breadth int) int {
	    return length * breadth
	  }

	 // function that returns an anonymous function
	func displayNumber() func() int {

	  number := 10
	  return func() int {
	    number++
	    return number
	  }
	}
	countNum := displayNumber()
	countNum() // 11
	countNum() // 12
	countNum() // 13
*/

/*
	golang closure: defined a func in a func
	// outer function
	func greet() func() string {

	  // local variable
	  name := "John"

	  // return a nested anonymous function
	  return func() string {
	    name = "Hi " + name
	    return name
	  }
	}
	message := greet()
	fmt.Println(message()) // "Hi John"


	func displayNumbers() func() int {
	  number := 0

	  // inner function
	  return func() int {
	  number++
	  return number
	  }
	}

	// returns a closure
	num1 := displayNumbers()

	fmt.Println(num1()) // 1
	fmt.Println(num1()) // 2
	fmt.Println(num1()) // 3

	// returns a new closure
	num2 := displayNumbers()
	fmt.Println(num2()) // 1
	fmt.Println(num2()) // 2
*/

// var varName string //<=> varName := ""
// assign(=) dung de gan value cho varName da declared
// underscore da duoc declared nen dung (=)
func namedReturn(a, b int) (result int, ok bool) {
	_ = a + b
	// passed result=0,ok=false to return
	return //o, false
	//return result, ok
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
	fmt.Println("lambda: ", lambda(3, 5))

	fmt.Println(math.Pow(3.5, 10))

	value := func() {
		fmt.Println("Welcome! to GeeksforGeeks")
	}
	value()

	tricks.Format("anonymous func")
	s1 := Student{
		firstName: "Naveen",
		lastName:  "Ramanathan",
		grade:     "A",
		country:   "India",
	}
	s2 := Student{
		firstName: "Samuel",
		lastName:  "Johnson",
		grade:     "B",
		country:   "USA",
	}

	// lambda := func(student Student) bool {
	//  code blabla
	// }
	students := []Student{s1, s2}
	students_filter := filter(students, func(student Student) bool {
		if student.grade == "B" {
			return true
		}
		return false
	})
	fmt.Println(students_filter)

	add := simple()
	fmt.Println(add(60, 7))

	tricks.Format("anonymous func()")

	// anonymous function = lambda
	my_name := "Cr7"
	func(name string) {
		fmt.Println("name =", my_name)
	}(my_name) // run lun

	tricks.Format("fibonacci func")
	fmt.Println("The factorial: ", factorial(5))

	tricks.Format("namedReturn")
	fmt.Println(namedReturn(4, 5))
	fmt.Println("loop() = ", loop())
	fmt.Println("loop2() = ", loop2())

}

func loop() float64 {
	ketqua := 100.0
	for i := 0; i < 5; i++ {
		if i == 4 {
			return ketqua
			// return //error, phai co named return, moi return kieu nay duoc
		} else {
			fmt.Println(i)
		}
	}

	// end func return phai co, bat buoc, vi declared co return float64
	return 500.0
}

func loop2() (result float64) {
	for i := 0; i < 5; i++ {
		if i == 4 {
			return
			// return //error, phai co named return, moi return kieu nay duoc
		} else {
			fmt.Println(i)
		}
	}
	// end func return phai co, bat buoc, vi declared co return float64
	return //passed ngam result
}