package main

import(
    "fmt"
    // "runtime"
    // "time"
    // "os"
    "strings"

    // "phong/tricks"
)

const NOTE string = `

Type:
    (func())
    (func(name string) string)
    (struct{})
    (struct{name string; old int})
    (interface{})
    (interface{one() string; two() int})  ...
    int float64 string (chan int) []string [2]int 

Expression: da declared, co gan value, thay co them ngoac nhon {} la expression
    (func() {return})
    (func(name string) string {return name})
    (struct{}{})
    (struct{name string; old int}{name:"thong", old:39})
    (interface{})
    (interface{one() string; two() int})  ...
    
    i := 3 // int(3)
    ch := make(chan int) // 


 `

func main() {
    repeat := strings.Repeat("-", 20)
    fmt.Println(repeat, "type Conversion", repeat)
    fmt.Println(repeat, "obj := (Type)(Expression)", repeat)
    // f := func(name string) string {return name}
    // var three = (func(name string) string)(func(name string) string {return name})
    var three = (func(name string) string {return name})
    // three := (func(name string) string {return name})
    fmt.Println("three", three("chi thong"))

    f := (func(name string) string{return string(name)})
    fmt.Println(f("chi thong"))


    fmt.Println(repeat, "obj_check, ok := (Expression).(Type)", repeat)
    fmt.Println(repeat, "type assertion", repeat)

    var h interface{}
    h = Human{Name:"Chi dung"}
    fmt.Println(h)

    human, ok := h.(interface{haha() string}) // struct Human co method haha() string ko
    fmt.Println(human, ok)

    var Person = (interface{})(struct{name string; old int}{name:"thong", old:38})
    fmt.Println(Person)

    ck, yes := Person.(struct{name string; old int})
    fmt.Println(ck, yes)


    var decorator Decorator
    decorator = Decorator(NewHuman)
    decorator.run("CHI THONG", 35)
    fmt.Println("decorator", decorator)

    one := &ONE{Name:"CHAU CHO"} // use pointer one update cho two
    two := TWO{ONE: one, Old:39}
    fmt.Println("two", two)
    fmt.Println("two.Name", two.Name)

    one.Name = "TIGER"
    fmt.Println("two.Name update via one", two.Name)

    one.Name = "LION"
    fmt.Println("two.Name update via one", two.Name)

}

type ONE struct {
    Name string
}

type TWO struct {
    *ONE
    Old int
}



type Human struct {
    Name string
    Old int
}

func(h Human) one() string{
    return h.Name
}

func(h Human) two() string{
    return h.Name
}

// non struct
type Decorator func(name string, old int) interface{}

func(d Decorator) run(name string , old int){
    repeat := strings.Repeat("-", 20)
    fmt.Println(repeat, "wrapped", repeat)
    fmt.Println(d(name, old))
    fmt.Println(repeat, "wrapped", repeat)
}

func NewHuman(name string, old int) interface{} {
    return &Human{Name:name, Old:old}
}

