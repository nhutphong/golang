package main

import(
    "fmt"
    // "runtime"
    // "time"
    // "os"
)

func main() {
    var h interface{}
    h = Human{Name:"Chi dung"}
    fmt.Println(h)

    human, ok := h.(interface{two() string})
    fmt.Println(human, ok)

    // f := func(name string) string {return name}
    var three = (func(name string) string)(func(name string) string {return name})
    fmt.Println("three", three("chi thong"))

    var Person = (interface{})(struct{name string; old int}{name:"thong", old:38})
    fmt.Println(Person)

    ck, yes := Person.(struct{name string; old int})
    fmt.Println(ck, yes)



}

type Human struct {
    Name string
}

func(h Human) one() string{
    return h.Name
}

func(h Human) two() string{
    return h.Name
}

