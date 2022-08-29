package main

import (
    "fmt"
    "phong/tricks"

)

func action (ch chan int, quit chan int) {
    tricks.Format("anonymous func()")

    for i := 0; i < 10; i++ {
        fmt.Println("go: ",<-ch)
    }

    fmt.Println("outside for {}")
    quit <- 0

}

func fibonacci(ch, quit chan int) {
    x, y, count := 0, 1 ,0
    tricks.Format("fibonacci()")

    dem := 1
    for {
        fmt.Printf("for fibonacci------------------------------------- start %v\n", dem)

        select {
        case ch <- x:
            x, y = y, x+y
            count++
            fmt.Println("case count: ", count)

        case <-quit:
            fmt.Println("quit")
            return
        }

        fmt.Printf("for fibonacci------------------------------------- end %v\n", dem)
        dem++
    }
}


func main() {
    ch := make(chan int)
    quit := make(chan int)
    go func() {
        tricks.Format("anonymous func()")

        for i := 1; i < 20; i++ {
            fmt.Printf("for %v -------------------------start\n", i)

            fmt.Println("go ch =  ",<-ch)

            fmt.Printf("for %v--------------------------end\n", i)
            fmt.Println()
        }

        quit <- 0
        fmt.Println("outside for {}")

    }()

    // go action(ch, quit)

    tricks.Format("main()")
    fibonacci(ch, quit)
}