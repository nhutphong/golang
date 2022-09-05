package main

import (  
    "fmt"
    "os"
    "time"
    // "log"
    "strings"
    // "bufio"
    "reflect"

    "phong/tricks"
)

func isError(err error) bool {
    if err != nil {
        fmt.Println(err.Error())
    }
    return (err != nil)
}

func ReadFile(name string) {
    slice, err := os.ReadFile(name)
    if isError(err) {
        return
    }

    for _, item := range slice {
        time.Sleep(time.Millisecond * 100)
        fmt.Println(item)
    }
    // os.Stdout.Write(slice) // write on terminal
}

// return phai la (lines []string) ; NOT []string
func arrLine(name string) (lines []string) {
    slice, err := os.ReadFile(name)
    if isError(err) {
        return
    }

    content := string(slice)
    // lines da declared named value
    lines = strings.Split(content, "\n")
    return lines
}


func main() {  
    fmt.Println()
    time.Sleep(time.Second)

    tricks.Format(`contents, err := os.ReadFile("test.txt")`)
    file_name := "test.txt"

    slice := []byte("thong")
    fmt.Println(slice)
    fmt.Println(string(slice))

    slice, err := os.ReadFile(file_name)
    fmt.Println(`reflect.TypeOf(slice) `, reflect.TypeOf(slice))
    fmt.Println(`reflect.TypeOf(slice) `, reflect.TypeOf(err))

    content := string(slice)
    fmt.Println("reflect.TypeOf(content) ", reflect.TypeOf(content))

    // lines := strings.Split(content, "\n")
    // fmt.Println("reflect.TypeOf(lines) ", reflect.TypeOf(lines))
    // // fmt.Println("lines ; ", lines)
    // fmt.Println("len(lines) ; ", len(lines))

    // for _, line := range lines {
    //     fmt.Println(line)
    //     time.Sleep(time.Second)
    // }


    for index, line := range arrLine(file_name) {
        fmt.Println("index ", index, " ", line)
        time.Sleep(time.Second)
    }





    tricks.Format(`contents, err := os.ReadFile("test.txt")`)
}