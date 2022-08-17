package main

import (
	"fmt"
)

/*
	channels nhu pipe = 1 var nhan va gui data cho goroutines dung no
	default main() la goroutines 1

	KENH := make(chan int, 10) // nhận, send data 10 lần 

	KENH <- 25      gan data cho KENH
	data := <- KENH    nhan data tu KENH

	value, ok := channel == maps
*/

func digits(number int, ch_digit chan int) {
	for number != 0 {
		digit := number % 10
		ch_digit <- digit
		number /= 10
	}
	close(ch_digit)
}

func calcSquares(number int, ch_square chan int) {
	fmt.Println("func calcSquares() start")
	sum := 0
	ch_digit := make(chan int)
	go digits(number, ch_digit)
	for digit := range ch_digit {
		sum += digit * digit
	}
	ch_square <- sum
	fmt.Println("func calcSquares() end")
}

func calcCubes(number int, ch_cube chan int) {
	fmt.Println("func calcCubes() start")
	sum := 0
	ch_digit := make(chan int)
	go digits(number, ch_digit)
	for digit := range ch_digit {
		sum += digit * digit * digit
	}
	ch_cube <- sum
	fmt.Println("func calcCubes() end")
}

func main() {
	number := 589
	ch_square := make(chan int)
	ch_cube := make(chan int)
	go calcSquares(number, ch_square)
	go calcCubes(number, ch_cube)
	// time.Sleep() cua main nen > nhat

	// vi trong goroutine calcSquares va calcCubes co write vao channel
	// ch_square va ch_cube nen thay phien nhau block goroutine main()
	// oonen 2 func co time scheduler va run

	// write ch_square <- data
	// read data <- ch_square

	fmt.Println("end goroutines channel")
	fmt.Println("end goroutines channel")
	fmt.Println("end goroutines channel")

	squares, cubes := <-ch_square, <-ch_cube
	fmt.Println("squares: ", squares)
	fmt.Println("cubes: ", cubes)
	fmt.Println("Final ouput", squares+cubes)
}
