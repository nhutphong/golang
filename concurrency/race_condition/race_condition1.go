package main

import (
    "fmt"
    "time"
    "sync"
)

const NOTE = ` 
    race conditions: vì nhiều goroutines khác như write cùng 1 lúc với var count nên phải dùng mutex
    or duplicate goroutines go write(); go write(); go write();
`

// var mutex = &sync.Mutex{}
var mutex sync.Mutex
const LOOP int = 10


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

// use global variable
var slice = []int{}
func read(ch chan int, lock chan bool) {
    // var list = []int{}
    print("\tPARENT read start\n")

    for i := 0; i < LOOP; i++ {
        fmt.Printf("goroutine read start %d\n", i)
        val := <-ch
        slice = append(slice, val)
        fmt.Println(val)
        fmt.Printf("goroutine read end %d\n", i)
    }

    lock <- true
    print("\tPARENT read end\n")
}

func main() {
    time.Sleep(time.Millisecond)

    ch := make(chan int)
    lock := make(chan bool)

    go write(ch)

    go read(ch, lock)

    <-lock //nếu ko có block chặn lại, thì dòng code dưới sẽ run trước 2 goroutines trên, slice=empty
    fmt.Println("slice", slice)
    // fmt.Println("list", <-dataCh)
    print("\t\tgoroutine/ main end\n")
}