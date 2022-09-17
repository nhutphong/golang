package main

import (  
    "fmt"
)


func main() {
    // defer func() {
    //     fmt.Println(recover()) // 3
    // }()

    // anonymous := func() string {
    //     return "anonymous"
    // }
    // fmt.Println(anonymous())


    // panic(3) // will replace panic 2
    // panic(2) // will replace panic 1
    // panic(1) // will replace panic 0
    // panic(0)

    sum(5, 2)
    fmt.Println("main() end")
}

const (
    NOTE = `use keyword defer, recover() chi nằm trong func or anonymous func, chỉ bắt được func,
        NOT goroutine  `
)


func catch() {
    fmt.Println("catch() start")
    if err := recover(); err != nil {
        fmt.Println(err)
    }
    fmt.Println("catch() end")
}

func sum(a int, b int) {  
    defer catch() //syntax: defer nameFunc() lun o tren cac func cua the xay ra panic, de bat no

    fmt.Printf("%d + %d = %d\n", a, b, a+b)
    divide(a, b) // error=panic ; jumpto func catch()
    fmt.Println("sum() end")
}

func divide(a int, b int) {  
    fmt.Printf("%d / %d = %d\n", a, b, a/b) // num / 0 ; error = panic
    fmt.Println("divide() end")
}