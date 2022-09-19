package main

import (
    "fmt"
    "sync"
)

var (
    balance int
)

func init() {
    balance = 100
}

func deposit(val int, wg *sync.WaitGroup, ch chan bool) {
    ch <- true
    balance += val
    <-ch
    wg.Done()
}

func withdraw(val int, wg *sync.WaitGroup, ch chan bool) {
    ch <- true
    balance -= val
    <-ch
    wg.Done()
}

func main() {
    var wg sync.WaitGroup
    ch := make(chan bool, 1)      // create the channel
    wg.Add(3)
    go deposit(20, &wg, ch)
    go withdraw(80, &wg, ch)
    go deposit(40, &wg, ch)
    wg.Wait()
    fmt.Printf("Balance is: %d\n", balance)
}