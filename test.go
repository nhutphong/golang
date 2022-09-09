package main

import (
    "fmt"
    "time"
)

func write(ch chan int) {
    fmt.Printf("\tGOROUTINE write START\n")

    for i := 1; i <= 5; i++ {
        ch <- i
        fmt.Println("write END", i, "to ch")
    }

    fmt.Printf("\tGOROUTINE write END\n")
}

func read(channel chan int) {
    fmt.Printf("\tGOROUTINE read START\n")

    for v := range channel {
        fmt.Println("read", v, "from channel")
        // time.Sleep(time.Second)
        // fmt.Printf("len %d  \n", len(channel))
        if len(channel) == 0 {
            time.Sleep(time.Second)
            // break
        }

    }

    fmt.Printf("\tGOROUTINE read END\n")
}


func main() {
    channel := make(chan int, 15)

    go write(channel)
    go write(channel)
    go write(channel)
    // go write(channel)

    go read(channel)
    

    time.Sleep(time.Second)
    fmt.Printf("\t\tGOROUTINE main() END\n")
    close(channel)
}
