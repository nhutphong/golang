package main

import (
    "fmt"
)

// channels nhu pipe = 1 var nhan va gui data cho goroutines dung no
// default main() la goroutines 1
// KENH := make(chan int)
// go func_name(name, kenh) =>  goroutines 2 ,...
// trong func_name se gan data cho KENH CD:              KENH <- 25
// number := <- kenh               tuc la number := 25 nhan tu KENH
// KENH <- 25      gan data cho KENH
// data := <- KENH    nhan data ut KENH


func digits(number int, dchnl chan int) {
    for number != 0 {
        digit := number % 10
        dchnl <- digit
        number /= 10
    }
    close(dchnl)
}

func calcSquares(number int, squareop chan int) {
    sum := 0
    dch := make(chan int)
    go digits(number, dch)
    for digit := range dch {
        sum += digit * digit
    }
    squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
    sum := 0
    dch := make(chan int)
    go digits(number, dch)
    for digit := range dch {
        sum += digit * digit * digit
    }
    cubeop <- sum
}

func main() {
    number := 589
    sqrch := make(chan int)
    cubech := make(chan int)
    go calcSquares(number, sqrch)
    go calcCubes(number, cubech)
    squares, cubes := <-sqrch, <-cubech
    fmt.Println("Final ouput", squares+cubes)
}