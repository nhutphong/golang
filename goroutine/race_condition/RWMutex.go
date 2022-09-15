package main

import (
    "sync"
    "time"
)

func main() {
    var i rwm

    go func() {
        for { // over loop
            i.inc() // free running counter
        }
    }()

    time.Sleep(1 * time.Second)
    println(i.read()) // sampling the counter
}

type rwm struct {
    sync.RWMutex
    i int
}

func (l *rwm) inc() {
    l.Lock()
    defer l.Unlock()
    l.i++
}
func (l *rwm) read() int {
    l.RLock()
    defer l.RUnlock()
    return l.i
}