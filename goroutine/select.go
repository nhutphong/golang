package main

import (
    "fmt"
    "time"
    "phong/tricks"

)

func action (ch chan int, quit chan bool) {
    tricks.Format("action() start")

    for i := 0; i < 10; i++ {
        fmt.Println("go: ",<-ch)
    }

    fmt.Println("action() end")
    quit <- true

}

func fibonacci(ch chan int, quit chan bool, result chan int) {
    x, y, count := 0, 1 ,0
    fmt.Println("\tfibonacci() start")

    dem := 1
    for {
        fmt.Printf("for fib start %v\n", dem)

        select {
        case ch <- y:
            x, y = y, x+y
            count++
            fmt.Println("case count: ", count)

        case <-quit:
            fmt.Println("quit")
            fmt.Println("\tfibonacci() end")
            result <- y
            return
        }

        fmt.Printf("for fib end %v\n", dem)
        dem++
    }

}

const (
    LOOP int = 12
)

func main() {
    ch := make(chan int)
    quit := make(chan bool)
    result := make(chan int)
    go func() {
        fmt.Println("\tREAD START")

        for i := 1; i < LOOP; i++ {
            fmt.Printf("for %v start\n", i)

            fmt.Println("go ch =  ",<-ch)

            fmt.Printf("for %v end\n", i)
            fmt.Println()
        }

        fmt.Println("\tREAD END")
        quit <- true

    }()

    // go action(ch, quit)
    go fibonacci(ch, quit, result)

    fmt.Println("result", <-result)
    time.Sleep(time.Second)
    tricks.Format("main() end")
}