package main

import (  
    "fmt"
)

func recovery() {  
    if r := recover(); r != nil {
        fmt.Println("recovered:", r)
    }
    // jumpto main()
}

func sum(a int, b int) {  
    defer recovery() //syntax: defer nameFunc() lun o tren cac func cua the xay ra panic, de bat no

    fmt.Printf("%d + %d = %d\n", a, b, a+b)

    done := make(chan bool)
    // go divide(a, b, done)
    divide(a, b, done) // error=panic ; jumpto func recovery()
    <-done
}

func divide(a int, b int, done chan bool) {  
    fmt.Printf("%d / %d = %d", a, b, a/b) // num / 0 ; error = panic
    done <- true

}

func main() {  
    sum(5, 0)
    fmt.Println("normally returned from main")
}