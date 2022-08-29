package main

import (
    "fmt"
    "time"
)

func main() {

    c1 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second) // chờ 2 second 
        c1 <- "result 1"
    }()

    select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(1 * time.Second): // sau 1 second sẽ run
        fmt.Println("chỉ chờ 1s, func run 2s, nên không chờ nữa")
    }


    c2 := make(chan string, 1)
    go func() {
        time.Sleep(2 * time.Second) // chờ 2 second 
        c2 <- "result 2"
    }()

    select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(3 * time.Second): // sau 3 second sẽ run 
        fmt.Println("timeout 2")
    }
}