package main

import (
	"fmt"
)

type Student struct {
	firstName string
	lastName  string
	grade     string
	country   string
}

func filter(students []Student, f func(Student) bool) []Student {
	var students_list []Student
	for _, student := range students {
		if f(student) == true {
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

func main() {
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
	f := filter(students, func(student Student) bool {
		if student.grade == "B" {
			return true
		}
		return false
	})
	fmt.Println(f)

	add := simple()
	fmt.Println(add(60, 7))

	my_name := "Cr7"
	func(name string) {
		fmt.Println("name =", my_name)
	}(my_name)
}
