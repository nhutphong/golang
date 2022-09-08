package main

import (
	"fmt"
	"log"
	"math"
	"phong/tricks"
)

/*
	struc chi chua fields = attributes
	interface chua methods

	human := struct{}{} // empty struct

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

/*
	thông thường muốn update value cho struct field, có thể access trực tiếp, NOT USE vì tính bảo mật,
	vì thế thể phải dùng func or method, bắt buộc phải pass pointer *struct vào để update value cho field

	func va method: p.FieldName = "update"  // rieng struct, ngâm hiểu có *
	cac types builtin in go: *int *string *float64, phai *varName = "value" // phai co * truoc varName
*/
func (p *Person) update(first, last string, age int) {
	p.FirstName = first
	p.LastName = last
	p.Age = age
}

// chi get value, NOT update field vì không pass pointer *Person 
func(p Person) info() {
	fmt.Println(p.FirstName, p.LastName, p.Age)
}

func update(p *Person, first, last string, age int) {
	p.FirstName = first
	p.LastName = last
	p.Age = age
}


type Car struct {
	Name, Model, Color string
	WeightInKg         float64
}

type Student struct {
	RollNumber int
	Name       string
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

// non-struct
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
    Address // embed struct xem thêm tại inheritance.go
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
	fmt.Printf("Person2: %#v\n", p2)

	// Uninitialized fields are set to their corresponding zero-value
	p3 := Person{FirstName: "Robert"}
	fmt.Println("Person3: ", p3)

	person := Person{
		FirstName: "John",
		LastName:  "Snow",
		Age:       45,
	}
	fmt.Println(person)
	person.update("dung", "thong", 35)
	fmt.Printf("%#v\n", person)

	update(&person, "Dog", "Cat", 100)
	fmt.Printf("%#v\n", person)
	//

	tricks.Format("NOT pointer")
	// Structs are value types.
	p5 := Point{10, 20}
	p4 := p5 // A copy of the struct `p1` is assigned to `p2`
	fmt.Println("p5 = ", p5)
	fmt.Println("p4 = ", p4)

	p4.X = 15
	fmt.Println("\nAfter modifying p2:")
	fmt.Println("p5 = ", p5)
	fmt.Println("p4 = ", p4)


	tricks.Format("so sánh pointer")
	p6 := Point{3.4, 5.2}
	p7 := Point{3.4, 5.2}
	if p6 == p7 {
		fmt.Println("Point p6 == p7 are equal.")
	} else {
		fmt.Println("Point p6 and p7 are not equal.")
	}


	tricks.Format("non-struct")
	// call non-struct dung (ngoặc tròn)
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	// type myInt int
	num1 := myInt(5)
    num2 := myInt(10)
    sum := num1.add(num2)
    fmt.Println("Sum is", sum)

    /*

    	// anonymous fields == embed struct == inheritance
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

    //
    tricks.Format("embed struct into struct = inheritance")
    human := Human{
        name: "Naveen",
        age:  50,
        Address: Address{
            city:  "Chicago",
            state: "Illinois",
        },
    }

    fmt.Println("human.name:", human.name)
    fmt.Println("human.age:", human.age)
    fmt.Println("human.Address.city:", human.Address.city)
    fmt.Println("human.Address.state:", human.Address.state)
	log.Println("vì embed struct nên có thể access fields=attributes")
    fmt.Println("human.city:", human.city)
    fmt.Println("human.state:", human.state)

}
