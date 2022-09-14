package main

import (
    "fmt"
    "time"
)

var k int

func say(lock chan bool, stop chan bool, s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s, k)
        lock<-true
        k++
        <-lock
    }
    stop<-true
}

func main() {
    lock := make(chan bool)
    stop := make(chan bool)
    go say(lock,stop, "world")
    go say(lock,stop, "hello")
    // say("hello")
    <-stop
    for i := 0; i < 10; i++ {
        
    }
    fmt.Println(k)
}