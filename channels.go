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


func digits(number int, ch_digit chan int) {
    for number != 0 {
        digit := number % 10
        ch_digit <- digit
        number /= 10
    }
    close(ch_digit)
}


func calcSquares(number int, ch_square chan int) {
    sum := 0
    ch_digit := make(chan int)
    go digits(number, ch_digit)
    for digit := range ch_digit {
        sum += digit * digit
    }
    ch_square <- sum
}


func calcCubes(number int, ch_cube chan int) {
    sum := 0
    ch_digit := make(chan int)
    go digits(number, ch_digit)
    for digit := range ch_digit {
        sum += digit * digit * digit
    }
    ch_cube <- sum
}


func main() {
    number := 589
    ch_square := make(chan int)
    ch_cube := make(chan int)
    go calcSquares(number, ch_square)
    go calcCubes(number, ch_cube)
    squares, cubes := <-ch_square, <-ch_cube
    fmt.Println("Final ouput", squares+cubes)
}