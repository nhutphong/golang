package main

import (
    "fmt"
    "sync"
    // "time"
)

type Data struct {
    total int
}

const (
    LOOP int = 10
)
func main() {
    c := 0
    n := 50

    ch := make(chan struct{}, n)
    chanWg := sync.WaitGroup{}
    chanWg.Add(1)
    go func() {
        for temp := range ch {
            c++
            fmt.Print(temp)
        }
        fmt.Println("count end")
        chanWg.Done()
    }()

    wg := sync.WaitGroup{}
    wg.Add(n)
    for i := 0; i < n; i++ {
        go func(i int) {
            ch <- struct{}{}
            fmt.Println("channel end", i+1)
            wg.Done()
        }(i)
    }
    wg.Wait()
    close(ch)

    chanWg.Wait()

    fmt.Println(c)

    // 200 = OK
}