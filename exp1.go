package main

import(
    "fmt"
    // "runtime"
    // "time"
    // "os"
)

func main() {
    p := Human{Name:"thong"}
    fmt.Printf("%#v\n", p)

    p2 := p.Pointer()
    p2.Name = "Dung"
    fmt.Printf("p2 %#v\n", p2)
    fmt.Printf("p1 %#v\n", p)

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

func (h *Human) Pointer() *Human {
    return h
}

