package main

import (
	"fmt"
	"math"
)

/*
	struc chi chua fields = attributes
	interface chua methods

	// anonymous struct
	emp3 := struct {
        firstName string
        lastName  string
        age       int
        salary    int
    }{
        firstName: "Andreah",
        lastName:  "Nikola",
        age:       31,
        salary:    5000,
    }

    fmt.Println("Employee 3", emp3)
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

// non-struct
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type myInt int

func (a myInt) add(b myInt) myInt {
	return a + b
}

/*
	// call: dung (ngoac tron, f is 1 argument, voi type=float64), ko phai la receiver=self gi
	// f := -math.Sqrt2
	my_float := MyFloat(-math.Sqrt2)
	fmt.Println(my_float.Abs())

*/


type Address struct {  
    city  string
    state string
}

type Human struct {
	// <name_fields> <type_field>
	
    name    string
    age     int
    address Address // nested struct
    /* Address 	// khi declared chi co Address=Type
    ngam hieu Address Address

    */
}

func format() {
	fmt.Println()
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println()
}


func main() {
	// call struct empty
	var p Person
	fmt.Println(p)

	// call struct co 3 fields dung {ngoac nhon}
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

	
	fmt.Println("(*ps).name is ", (*ps).Name)
	fmt.Println("ps.name is ",ps.Name)

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

	format()

	// call non-struct dung (ngoặc tròn)
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	format()

	// type myInt int
	num1 := myInt(5)
    num2 := myInt(10)
    sum := num1.add(num2)
    fmt.Println("Sum is", sum)

    format()

    /*

    	// anonymous fields
		type Person struct {  
		    string
		    int
		}	

	    person := Person{
	        string: "naveen",
	        int:    50,
	    }
	    fmt.Println(person.string)
	    fmt.Println(person.int)
    */

    format()

    //
    human := Human{
        name: "Naveen",
        age:  50,
        address: Address{
            city:  "Chicago",
            state: "Illinois",
        },
    }

    fmt.Println("Name:", human.name)
    fmt.Println("Age:", human.age)
    fmt.Println("City:", human.address.city)
    fmt.Println("State:", human.address.state)
}
