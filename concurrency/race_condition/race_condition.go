package main

import (
    "fmt"
    "time"
    "sync"
)

// var mutex = &sync.Mutex{}
var mutex sync.Mutex
const LOOP = 5


// race conditions: vì nhiều goroutines write cùng 1 lúc với var count nên phải dùng mutex
func write(ch chan int) {
    var count int = 1
    print("\tPARENT write start\n")
    for i := 0; i < LOOP; i++ {
        go func(ch chan int, index int) {
            fmt.Println("anonymous write start ", index)

            mutex.Lock()
            count+=count
            ch <- count
            mutex.Unlock()

            fmt.Println("anonymous write end", index)
        }(ch, i)
    }
    print("\tPARENT write end\n")
}


func read(ch chan int, dataCh chan []int) {
    var list = []int{}
    print("\tPARENT read start\n")

    for i := 0; i < LOOP; i++ {
        fmt.Printf("goroutine read start %d\n", i)
        val := <-ch
        list = append(list, val)
        fmt.Println(val)
        fmt.Printf("goroutine read end %d\n", i)
    }
    
    dataCh <- list
    print("\tPARENT read end\n")
}

func main() {
    time.Sleep(time.Millisecond)

    ch := make(chan int)
    dataCh := make(chan []int)

    go write(ch)
    go read(ch, dataCh)

    
    slice := <-dataCh
    for _, val := range slice {
        fmt.Println("val", val)
    }
    
    // fmt.Println("list", <-dataCh)
    // time.Sleep(time.Second)
    print("\t\tgoroutine/ main end\n")
}