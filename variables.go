package main

import "fmt"

func main() {

	// fmt.Printf("%s is %d years old", name, age) //print cosole 
	// fmt.Println("%s is %d years old", name, age) //print console xuong dong
	// res := fmt.Sprintf("%s is %d years old", name, age) // return res string

	// Declaring Variables
	var myStr string = "Hello"
	var myInt int = 100
	var myFloat float64 = 45.12
	fmt.Println(myStr, myInt, myFloat)

	//================================

	// Multiple Declarations
	var (
		employeeId          int    = 5
		firstName, lastName string = "Satoshi", "Nakamoto"
	)
	fmt.Println(employeeId, firstName, lastName)

	//================================

	// Short variable declaration syntax
	name := "Rajeev Singh"
	age, salary, isProgrammer := 35, 50000.0, true

	fmt.Println(name, age, salary, isProgrammer)


	// Type inference
	var myname = "Rajeev Singh" // Type declaration is optional here.
	fmt.Printf("Variable 'myname' is of type %T\n", myname)

	//================================

	// Multiple variable declarations with inferred types
	var ho, ten, tuoi, luong = "John", "Maxwell", 28, 50000.0

	fmt.Printf("ho: %T, ten: %T, tuoi: %T, luong: %T\n",
		ho, ten, tuoi, luong)


	var (
		dau, cuoi string
		old                 int
		total              float64
		isConfirmed         bool
	)

	fmt.Printf("dau: %s, cuoi: %s, old: %d, total: %f, isConfirmed: %t\n",
		dau, cuoi, old, total, isConfirmed)
}