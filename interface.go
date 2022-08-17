package main

import "fmt"

// Định nghĩa interface I với 1 method là M() và N()
// interface noi chua cac methods
type AbstractHuman interface {
	get_name() string
	get_age() int
	// get_city(city string) string // co (arg_city) return string
}

// Định nghĩa struct T = class
type Human struct {
	name string
	age  int
}

// implement interface
func (self Human) get_name() string {
	return self.name
}

// implement interface
func (self Human) get_age() int {
	return self.age
}

// var human *Human = &Human{name: "dung", age:28} //cach2
// var human        = &Human{name: "dung", age:28} //default cach1

func change_age_human(human *Human, age int) int {
	human.age = age
	return human.get_age()
}

// Hàm describe có tham số truyền vào là một empty interface
// do đó khi thực thi ta có thể truyền vào kiểu dữ liệu nào cũng được
func describe(auto interface{}) {
	fmt.Printf("value: %v, type: %T\n", auto, auto)
}

func main() {
	// var human Human = Human{name: "dung", age:28}
	// var abstract AbstractHuman = Human{name: "dung", age:28}

	var human = &Human{name: "dung", age: 28}
	fmt.Println(human.get_name())
	fmt.Println(human.get_age())
	fmt.Println("change_age_human: ", change_age_human(human, 35))

	fmt.Println(human.age)

	var auto interface{} // auto la value tuy y -> int string array slice map
	describe(auto)

	auto = 27
	describe(auto)

	auto = "chithong"
	describe(auto)

	auto = [...]int{}
	describe(auto)

	auto = make([]int, 5)
	describe(auto)

	auto = make(map[string]Human)
	describe(auto)
}
