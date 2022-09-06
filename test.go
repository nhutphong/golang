package main

import (
    "fmt"
    "time"
)

func main() {

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
