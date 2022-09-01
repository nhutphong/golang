package main

import (  
    "fmt"
    "os"
)

func main() {  
    contents, err := os.ReadFile("test.txt")
    if err != nil {
        fmt.Println("File reading error", err)
        return
    }
    fmt.Println("Contents of file:", string(contents))
}