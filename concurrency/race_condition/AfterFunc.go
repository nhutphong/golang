package main

import (
    "fmt"
    "time"
)

func main() {
    oneSecond := 1 * time.Second
    i := uint64(0)
    ch := make(chan struct{})

    time.AfterFunc(oneSecond, func() {
        close(ch)
    })

loop:
    for {
        select {
        case <-ch:
            break loop
        default:
            i++
        }
    }

    fmt.Println(i) // 104_178_671
}
