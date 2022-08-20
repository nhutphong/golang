package main

import (
	"fmt"
	"time"
	"phong/tricks"
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

func process(ch chan string) {
	time.Sleep(5000 * time.Millisecond)
	ch <- "process successful"
}

func server1(ch chan string) {  
    ch <- "from server1"
}
func server2(ch chan string) {  
    ch <- "from server2"

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

	squares, cubes := <-ch_square, <-ch_cube
	fmt.Println("squares: ", squares)
	fmt.Println("cubes: ", cubes)
	fmt.Println("Final ouput", squares+cubes)

	tricks.Format()

	/*
		select { case expression } == match case python
		thuong dung voi goroutine channel
	*/

	/* Deadlock and default case

		func main() {  
		    ch := make(chan string)
		    select {
			case <-ch:
			default:
				fmt.Println("co default se ko co deadlock")
		    }

		    select {} // chi co 1 minh select cung deadlock
		}
	*/

	output1 := make(chan string)
    output2 := make(chan string)
    go server1(output1)
    go server2(output2)
    time.Sleep(1 * time.Second)
    select {
    case s1 := <-output1:
        fmt.Println(s1)
    case s2 := <-output2:
        fmt.Println(s2)
    }

    tricks.Format()

    ch := make(chan string)
	go process(ch)
	for {
		time.Sleep(500 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println("received value: ", v)
			return
		default:
			fmt.Println("no value received")
		}
	}

}
