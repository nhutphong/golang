package main

import (
    "fmt"
    "strings"
)

type Animal interface {
    Walk()
    ChangeAge(Age int)
}

type Dog struct {
    Age int
}


// implement Animal, binh thuong khong co pointer receiver
func (d Dog) Walk() {
    fmt.Println("Dog Walk Age", d.Age)
}

// implement Animal, nhung co pointer receiver
func (d *Dog) ChangeAge(age int) {
    d.Age = age
}

type Cat struct {
    // Dog // embed  struct
    *Dog // embed pointer struct; access methods() va fields; su dung pointer dog, update
    // ben ngoai khong can thong qua cat


    // Animal // embed interface ; chi access methods() NOT fields; struct pass vao phai la pointer

    // *Animal // embed interface; KHONG THE NAO EMBED POINTER INTERFACE
}





func main() {
    repeat := strings.Repeat("-", 20)
    fmt.Println(repeat, "access via Animal interface", repeat)

    var animal Animal
    dog2 := &Dog{Age: 30}
    animal = dog2
    animal.Walk()

    animal.ChangeAge(300)
    animal.Walk()



    fmt.Println(repeat, "access via Cat struct", repeat)
    dog := &Dog{Age:30}
    // cat := Cat{Dog: dog}
    cat := Cat{Dog: dog}

    fmt.Printf("%#v\n", dog)
    fmt.Printf("%#v\n", cat)
    cat.Walk()

    cat.ChangeAge(1000)
    cat.Walk()
    fmt.Println("cat.Dog", cat.Dog)
    fmt.Println("cat.Dog", *cat.Dog)
    // fmt.Println(cat.Age)

    // animal = dog
    // animal.Walk()

    // dog.Age = 580
    // animal.Walk()

    // animal.ChangeAge(1000)
    // animal.Walk()

    fmt.Println(repeat, NOTE, repeat)
}

const NOTE string = `
    
tránh lỗi dùng Animal interface, access method() có pointer receiver,thì cứ create struct là
pointer struct la được 

 `