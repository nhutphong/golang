package main

import (
    "fmt"
)


func main() {   
    fmt.Println()

    for i := 0; i < 5; i++ {
        defer fmt.Println(i)
    }

    for i := 'a'; i < 'f'; i++ {
        defer fmt.Printf("%c\n", i)   
    }

    fmt.Println("main() end")
}

const (
    NOTE = `defer: chỉ run sau return(kể cả return của block scopes trong func)
        vào sau ra trước = last in first out(LIFO) = stack(ngăn xếp)


        FIFO = queue(hàng đợi) = vào trước ra trước
        use package: "container/list"
    `
)
