package main

import (
    "fmt"
    // "time"
)


func main() {
   h := Human{Name:"dung"}
   run(h)
}

func run(ok OK) {
    ok.Show()
}


type OK interface {
    Show()
}

type Human struct {
    Name string
}

func (h Human) Show() {
    fmt.Println(h.Name)
}