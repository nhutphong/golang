package main

import (
	"fmt"
	"strings"
)

type Animal interface {
	Walk() string
	ChangeAge(age int)
	ChangeName(name string)
}

type Dog struct {
	Name string
	Age  int
}

// implement Animal, binh thuong khong co pointer receiver
func (d Dog) Walk() string {
	return fmt.Sprintf("{Name: %s, Age: %d}", d.Name, d.Age)
}

// implement Animal, nhung co pointer receiver
func (d *Dog) ChangeAge(age int) {
	d.Age = age
}

func (d *Dog) ChangeName(name string) {
	d.Name = name
}

type Cat struct {
	// Dog // embed  struct
	*Dog // embed pointer struct; access methods() va fields; su dung pointer dog, update
	// ben ngoai khong can thong qua cat

	// Animal // embed interface ; chi access methods() NOT fields;

	// *Animal // embed interface; KHONG THE NAO EMBED POINTER INTERFACE
}

func main() {
	repeat := strings.Repeat("-", 20)
	fmt.Println(repeat, "access via interface-Animal", repeat)

	/*
		interface chi access methods(), NOT access fields of struct
	*/
	var animal Animal
	dog2 := &Dog{Name: "thong", Age: 30} //pointer
	animal = dog2                        //pointer
	println("animal = ", animal)
	fmt.Println("animal.walk() = ", animal.Walk())

	println()
	println("update dog via interface-Animal")
	println("interface only access methods(), NOT fields")
	println("animal.ChangeAge(300)")
	println("animal.ChangeName(\"nguyen thong dung\")")
	//animal.Name = "nguyen chi thong" // NOT access (no fields)
	//animal.Age = 300 // NOT access (no fields)
	animal.ChangeAge(300)                  // animal is pointer moi access duoc ChangeAge(300)
	animal.ChangeName("nguyen thong dung") // animal is pointer moi access duoc ChangeName("nguyen thong dung")
	fmt.Println("animal.walk() = ", animal.Walk())

	println()
	fmt.Println(repeat, "access via struct-cat", repeat)
	dog := &Dog{Name: "dung", Age: 30} //pointer
	cat := Cat{Dog: dog}               //pointer

	fmt.Printf("dog = %#v\n", dog)
	fmt.Printf("cat = %#v\n", cat)
	fmt.Println("cat.walk() = ", cat.Walk())
	println()

	println("update dog via struct-cat")
	cat.ChangeAge(1000)
	cat.ChangeName("NGUYEN DUNG THONG")
	//cat.Age = 1000 //OK
	//cat.Name = "NGUYEN DUNG THONG" //OK
	fmt.Println("cat.walk() = ", cat.Walk())
	fmt.Println("cat.Dog = ", cat.Dog)
	fmt.Println("*cat.Dog = ", *cat.Dog)
	fmt.Println("dog = ", dog)
	println()

	println("update struct-cat via dog")
	dog.Name = "thong dung nguyen"
	dog.Age = 5000
	fmt.Println("dog = ", dog)
	fmt.Println("cat.walk() = ", cat.Walk())
	fmt.Println("cat.Dog = ", cat.Dog)
	fmt.Println("*cat.Dog = ", *cat.Dog)

	fmt.Println(repeat, NOTE, repeat)
}

const NOTE string = `
tránh lỗi dùng Animal interface, access method() có pointer receiver, thi interface phai duoc cho 1 struct-pointer, NOT struct, vd:
	var animal Animal // interface
	dog := &Dog{Name: "thong", Age: 35} // pointer
	animal = dog // interface animal duoc gan struct-pointer dog

	tranh panic: cu gan struct-pointer cho interface, du implement interface co hay ko co pointer receiver cung duoc
 `
