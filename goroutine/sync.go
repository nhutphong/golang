package main

import (
    "fmt"
    "time"
    "sync"
)

var wg sync.WaitGroup

func write(c chan int, someValue int) {
    defer wg.Done()
    c <- someValue * 5
    fmt.Println("GOROUTINE write() END")
}

func read(c chan int) {

}

func main() {
    buffer := make(chan int,10)
    for i := 0; i < 10; i++ {
        wg.Add(1)
        go write(buffer, i)
    }

    wg.Wait()
    close(buffer)
    for item := range buffer {
        fmt.Println(item)
    }

    time.Sleep(time.Second)
    fmt.Println("GOROUTINE main() END")
}