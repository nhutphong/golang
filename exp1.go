package main

import (
    "fmt"
    // "time"
    // "sync"
)

func write(ch chan<- int, i int, total int) {
    fmt.Printf("\tGOROUTINE write START\n")
    i++

    for count := 0; count < total; count++ {
        ch <- i
    }

    fmt.Printf("\tGOROUTINE write END\n")
}

func read(ch <-chan int, resch chan<- int, total int) {
    fmt.Printf("\tGOROUTINE read START\n")
    i := <-ch
    i *= 10

    for count := 0; count < total; count++ {
        resch <- i
        i *= 10
    }

    fmt.Printf("\tGOROUTINE read START\n")
}

func main() {
    chw := make(chan int)
    chr := make(chan int, 16)
    // defer close(chw)
    // defer close(chr)
    // defer fmt.Println("SuperGoroutineControl() END")
    // defer time.Sleep(time.Second)

    go write(chw,4, 3)
    go write(chw,29, 4)
    // go write(chw,19, 3)
    // go write(chw,29, 3)

    // go read(chw, chr, 3)
    go read(chw, chr, 7)

    // for i := 0; i < 16; i++ {
    //     fmt.Println(<-chr)
    // }

    for res := range chr {
        fmt.Println(res)
        if len(chr) == 0 {
            break
        }
    }

    // time.Sleep(time.Second)
    fmt.Println("goroutine main")

}
