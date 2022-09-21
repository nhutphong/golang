package main

import(
    "fmt"
    "runtime"
    // "time"
    // "os"
)

func main() {
    c := make(chan int)
    go func() {
        defer func() {c <- 1}()
        defer fmt.Println("Go")
        func() {
            defer fmt.Println("C")
            runtime.Goexit() // stop programing
        }()
        // not run
        fmt.Println("Java")
    }()
    <-c
}