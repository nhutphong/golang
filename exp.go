package main

import "fmt"

func main() {
    func() {
        fmt.Println(recover()) // 3
    }()
    anonymous := func() string {
        return "anonymous"
    }()

    panic(3) // will replace panic 2
    panic(2) // will replace panic 1
    panic(1) // will replace panic 0
    panic(0)
}