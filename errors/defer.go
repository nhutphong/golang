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
    fmt.Println("foo() :", foo())
}

func foo() (result string) {
    defer func() {
        result = "Change World" // last change named result of return 
    }()
    return "Hello World"
}

// close a file
func CopyFile(dstName, srcName string) (written int64, err error) {
    src, err := os.Open(srcName)
    if err != nil {
        return
    }
    defer src.Close()

    dst, err := os.Create(dstName)
    if err != nil {
        return
    }
    defer dst.Close()

    return io.Copy(dst, src)
}

const (
    NOTE = `defer: chỉ run sau return(kể cả return của block scopes trong func)
        vào sau ra trước = last in first out(LIFO) = stack(ngăn xếp)


        FIFO = queue(hàng đợi) = vào trước ra trước
        use package: "container/list"
    `
)
