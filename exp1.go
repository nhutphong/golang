package main

import(
    "fmt"
    // "runtime"
    // "time"
    // "os"
)

func main() {
    var i interface{}
    var ch chan int
    var p *int
    fmt.Println(i, ch, p)

    fmt.Printf("interface{} = %#v\n",i)
    fmt.Printf("interface{} = %T\n",i)
    println("interface{} = ",i)

    fmt.Printf("ch = %#v\n",ch)
    fmt.Printf("pointer = %#v\n",p)

    var f func()
    var slice []int
    var m map[string]string
    fmt.Println(f, slice, m)

    fmt.Printf("func() = %#v\n",f)
    fmt.Printf("slice = %#v\n",slice)
    fmt.Printf("map = %#v\n",m)

    // ha := (func())(nil)
    fmt.Println("(fun())(nil) == nil", (func())(nil) == nil)

}