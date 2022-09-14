package main

import (
	"fmt"
	// "log"
	"time"
	
	"phong/tricks"
)


func goroutine(unbuffer chan int) {
	//run goroutine thu 2
	go func() { 
		for i:=0; i<10; i++ {
			fmt.Println(<-unbuffer)
		}
	}()

	unbuffer <- 1
	unbuffer <- 2
	unbuffer <- 3
	unbuffer <- 4
	unbuffer <- 5
}


func write(unbuffer chan int, name string) {
	fmt.Printf("GOROUTINE write() START %s\n", name)

	for i := 1; ; i++ {
		tricks.FormatTwo("WRITE FOR START", i)
		unbuffer <- i
		tricks.FormatTwo("WRITE FOR END", i)
	}

	// close(unbuffer)
	fmt.Printf("GOROUTINE write() END %s\n", name)
	fmt.Println()
}

func read(unbuffer chan int, endch chan bool, total int, name string) {
	fmt.Printf("GOROUTINE read() START %s\n", name)

	for i := 1; i <= total; i++ {
		tricks.FormatTwo("READ FOR START", i)
		fmt.Println("read value", <-unbuffer, "from ch")
		tricks.FormatTwo("READ FOR END", i)
	}

	endch<-true
	fmt.Printf("GOROUTINE read() END %s\n", name)
}

func SuperGoroutineControl() {

	unbuffer := make(chan int,5) // STEP 2
	// endch := make(chan bool)
	// defer close(unbuffer)
	// defer fmt.Println("SuperGoroutineControl() END")
	// defer time.Sleep(time.Second)

	go write(unbuffer, "ONE")
	go write(unbuffer, "TWO")
	go write(unbuffer, "THREE")
	go write(unbuffer, "FOUR")
	go write(unbuffer, "FIVE")

	// go read(unbuffer, endch, 5, "ONE")
	for i := 0; i < 5; i++ {
		fmt.Println("main read ",<-unbuffer)
	}

	time.Sleep(time.Second) // phai co de jumpto goroutine read || write, ; se end sau cung  
	// <-endch
	fmt.Println("SuperGoroutineControl() END")
	// close(unbuffer)
}


//STEP 1
func main() {


	//hien tai co goroutine main
	tricks.Format("unbuffered-channel")
	tricks.Format("START GOROUTINE MAIN()")

	// unbuffer := make(chan int) // STEP 2
	// // defer close(unbuffer)
	// // defer time.Sleep(time.Second)

	// go write(unbuffer, "ONE")
	// go write(unbuffer, "TWO")
	// go write(unbuffer, "THREE")
	// go write(unbuffer, "FOUR")
	// go write(unbuffer, "FIVE")

	// go read(unbuffer, 15, "ONE")

	// go read(unbuffer, 5, "TWO")
	// go read(unbuffer, 5, "THREE")

	// go readUnBufferUseRange(unbuffer)

	SuperGoroutineControl()




	

	time.Sleep(time.Second)
	tricks.Format("END GOROUTINE MAIN()")
	// close(unbuffer)
}