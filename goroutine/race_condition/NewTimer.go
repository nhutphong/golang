package main

import "time"

func main() {
    ch := make(chan int)
    go func() {
        timeout := time.NewTimer(1 * time.Second)
        defer timeout.Stop()
        i := 1
        for {
            select {
            case <-timeout.C:
                ch <- i
                return
            default:
                i++
            }
        }
    }()

    println(<-ch)
}