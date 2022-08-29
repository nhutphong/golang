package main

import (
	"fmt"
	"time"
	"phong/tricks"
)

/*
	1 func = goroutine chi nen lam 1 viec send or receive to channel, 
	neu lam 2 viec send va receive to channel se error deadlock

	khi dung channel se khong can dung var wg WaitGroup, de stop main() tạm thời, chờ các goroutine run
	xong, main() mới run tiếp

	khi 1 goroutine send=write data to 1 channel, thì phải có 1 goroutine nào đó receive=read data từ
	channel đó, nếu không sẽ dẫn đến deadlock
	NOTE: channel được send data thì phải receive, or receive phải send

	KENH := make(chan int) // write, read on channel bao nhiêu lần cũng được 
	KENH := make(chan int, 10) // nhận, send data 10 lần 

	KENH <- 25      	send=write data=25 cho KENH
	data := <- KENH 	receive=read data=25 tu KENH, func main() tạm dừng chờ read tata from channel 

	value, ok := channel == maps
*/

func digits(number int, ch_digit chan int) {
	for number != 0 {
		digit := number % 10
		ch_digit <- digit // send to channel
		number /= 10
	}
	close(ch_digit) // close channel
}

func calcSquares(number int, ch_square chan int) {
	fmt.Println("func calcSquares() start")
	sum := 0
	ch_digit := make(chan int)
	go digits(number, ch_digit)
	for digit := range ch_digit {
		sum += digit * digit
	}
	ch_square <- sum // write to channel
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
	ch_cube <- sum // write to channel
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

func printNum() {
	for i := 1; i <= 10; i++ {
		time.Sleep(150 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func printChar() {
	for i := 'a'; i <= 'a' + 26; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

// "Z"
func printCharUpper() {
	for i := 'A'; i <= 'A' + 26; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}


func main() {

	tricks.Format("GOROUTINES THONG START")
	number := 589
	ch_square := make(chan int)
	ch_cube := make(chan int)
	go calcSquares(number, ch_square)
	go calcCubes(number, ch_cube)
	// time.Sleep() cua main nen > nhat

	// vi trong goroutine calcSquares va calcCubes co write vao channel
	// ch_square va ch_cube nen thay phien nhau block goroutine main()
	// oonen 2 func co time scheduler va run

	// ch_square <- data //send data to ch_square
	// data <- ch_square  // receiver data from ch_square

	squares, cubes := <-ch_square, <-ch_cube // main() sẽ wait, để read data từ channel rồi mới run tiếp 
	fmt.Println("squares: ", squares)
	fmt.Println("cubes: ", cubes)
	fmt.Println("Final ouput", squares+cubes)

	tricks.Format("GOROUTINES THONG END")

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

	tricks.Format("GOROUTINES DUNG START")
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
	tricks.Format("GOROUTINES DUNG END")


	tricks.Format("GOROUTINES PHONG START")
    ch := make(chan string)
	go process(ch)
	for {
		time.Sleep(500 * time.Millisecond)
		select {
		case value := <-ch:
			fmt.Println("received value: ", value)
			return // end forever loop ; end main()
			// break
		default:
			fmt.Println("no value received")
		}
	}
	tricks.Format("GOROUTINES PHONG END")

	fmt.Println("END MAIN")

}
