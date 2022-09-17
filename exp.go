package main

import "fmt"


func main() {
    defer func() {
        fmt.Println(recover()) // 3
    }()

    anonymous := func() string {
        return "anonymous"
    }
    fmt.Println(anonymous())


    panic(3) // will replace panic 2
    panic(2) // will replace panic 1
    panic(1) // will replace panic 0
    panic(0)
}

func catch() {
    if err := recover(); err != nil {
        fmt.Println(err)
    }
    fmt.Println("catch() end")
}