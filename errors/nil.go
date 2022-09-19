package main

import (
    "fmt"
)

func main() {
    var m map[string]int // nil
    delete(m, "foo")

    var x []int // nil
    a := x[:]
    b := x[0:0]
    c := x[:0:0]
    // Print three "true".
    fmt.Println(a == nil, b == nil, c == nil)


    var s []int // nil
    for range s {
    }

    var m1 map[string]int // nil
    for range m1 {
    }

    // vi co underscore nen ok, NOT error
    var a *[5]int // nil
    for i, _ := range a {
        fmt.Print(i)
    }
}