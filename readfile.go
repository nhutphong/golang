package main

import (  
    "fmt"
    "os"
    "time"
    // "log"
    // "strings"
    "bufio"

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




func main() {  
    fmt.Println()
    tricks.Format(`contents, err := os.ReadFile("test.txt")`)
    file_name := "test.txt"
    // content := ReadFileConvertToContent(file_name)
    // fmt.Println(content)

    slice := []byte("thong")
    fmt.Println(slice)
    fmt.Println(string(slice))


    // filePath := os.Args[1]
    readFile, err := os.Open(file_name)
  
    if err != nil {
        fmt.Println(err)
    }
    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    var fileLines []string
  
    for fileScanner.Scan() {
        fileLines = append(fileLines, fileScanner.Text())
    }
  
    readFile.Close()
  
    for _, line := range fileLines {
        fmt.Println(line)
    }
     
    fmt.Println(fileLines)


    tricks.Format(`contents, err := os.ReadFile("test.txt")`)
}