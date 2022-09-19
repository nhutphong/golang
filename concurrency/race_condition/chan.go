package main

import (
    "time"
)

func main() {
    ch := make(chan int, 1)
    ch <- 1
    timeout := time.NewTimer(1 * time.Second)
loop:
    for {
        select {
        case <-timeout.C:
            timeout.Stop()
            break loop
        default:
            ch <- 1 + <-ch
        }
    }

    println(<-ch)
}