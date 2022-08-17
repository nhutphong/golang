package main

import (
	"fmt"
)

/*
	struc chi chua fields = attributes
	interface chua methods
*/
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Car struct {
	Name, Model, Color string
	WeightInKg         float64
}

type Student struct {
	RollNumber int
	Name       string
}

type Employee struct {
	Id   int
	Name string
}

type Point struct {
	X float64
	Y float64
}

func format() {
	fmt.Println()
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println()
}


func main() {
	// Declaring a variable of a `struct` type
	var p Person
	fmt.Println(p)

	// Declaring and initializing a struct using a struct literal
	p1 := Person{"Rajeev", "Singh", 26}
	fmt.Println("Person1: ", p1)

	// Naming fields while initializing a struct
	p2 := Person{
		FirstName: "John",
		LastName:  "Snow",
		Age:       45,
	}
	fmt.Println("Person2: ", p2)

	// Uninitialized fields are set to their corresponding zero-value
	p3 := Person{FirstName: "Robert"}
	fmt.Println("Person3: ", p3)

	format()

	c := Car{
		Name:       "Ferrari",
		Model:      "GTC4",
		Color:      "Red",
		WeightInKg: 1920,
	}

	// Accessing struct fields using the dot operator
	fmt.Println("Car Name: ", c.Name)
	fmt.Println("Car Color: ", c.Color)

	// Assigning a new value to a struct field
	c.Color = "Black"
	fmt.Println("Car: ", c)


	format()
	// instance of student struct type
	s := Student{11, "Jack"}

	// Pointer to the student struct
	ps := &s
	fmt.Println("ps: ",ps)

	// Accessing struct fields via pointer
	fmt.Println("(*ps).name is ", (*ps).Name)
	fmt.Println("ps.name is ",ps.Name) // Same as above: No need to explicitly dereference the pointer

	ps.RollNumber = 31
	fmt.Println(ps)
	//

	format()

	pEmp := new(Employee)

	pEmp.Id = 1000
	pEmp.Name = "Sachin"

	fmt.Println(pEmp)


	format()
	// Structs are value types.
	p5 := Point{10, 20}
	p4 := p5 // A copy of the struct `p1` is assigned to `p2`
	fmt.Println("p5 = ", p5)
	fmt.Println("p4 = ", p4)

	p4.X = 15
	fmt.Println("\nAfter modifying p2:")
	fmt.Println("p5 = ", p5)
	fmt.Println("p4 = ", p4)

	format()
	
	p6 := Point{3.4, 5.2}
	p7 := Point{3.4, 5.2}

	if p6 == p7 {
		fmt.Println("Point p6 and p7 are equal.")
	} else {
		fmt.Println("Point p6 and p7 are not equal.")
	}
}
