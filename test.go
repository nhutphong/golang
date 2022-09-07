package main

import (
    "fmt"
    "time"
)



func main() {

	animal := Animal{"dog", 35}
	dog := Dog{animal, "red"}

	fmt.Println(animal)
	fmt.Printf("%#v\n", dog)
	fmt.Println(dog.Name, dog.age, dog.Color)

	var dict map[bool]bool
    fmt.Print("Default Zero Value of a map: ")
    fmt.Println(dict)

    var channel chan int 
    fmt.Print("Default Zero Value of a channel: ")
    fmt.Println(channel)

    var face interface{}
    fmt.Print("Default Zero Value of a interface: ")
    fmt.Println(face)

    var slice []int
    fmt.Println("slice == nil", slice == nil)
    fmt.Print("Default Zero Value of a slice: ")
    fmt.Println(slice)

    var f func()
    fmt.Print("Default Zero Value of a func: ")
    fmt.Println(f)

    var pointer *int
    fmt.Print("Default Zero Value of a pointer: ")
    fmt.Println(pointer)


}

type Animal struct {
	Name string
	age int
}

type Dog struct {
	Animal
	Color string
}


func Ticker () {

    ticker := time.NewTicker(500 * time.Millisecond)
    unbuffer := make(chan bool)

    go func() {
        for {
            select {
            case <-unbuffer:
            	fmt.Println("end timer")
                return
            case t := <-ticker.C:
                fmt.Println("Tick at", t)
            }
        }
    }()

    fmt.Println("one")
    time.Sleep(1600 * time.Millisecond)
    fmt.Println("two")
    ticker.Stop()
    fmt.Println("three")
    unbuffer <- true
    fmt.Println("Ticker stopped")
}