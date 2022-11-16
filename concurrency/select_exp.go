package main

import (
	"fmt"
	"time"
)

func main() {

	number := make(chan int)
	message := make(chan string)

	for i := 0; i < 10; i++ {
		_ = i
		go channelNumber(number)
		go channelMessage(message)

	}

	for i := 0; i < 10; i++ {
		_ = i
		loop(number, message)

	}

}

func loop(number chan int, message chan string) {
	select {
	case firstChannel := <-number: //block, await goroutine write data vao number<- 15
		fmt.Println("Channel Data:", firstChannel)

	case secondChannel := <-message: //block, await goroutine write data vao message<- "hahah"
		fmt.Println("Channel Data:", secondChannel)
	}
}

func channelNumber(number chan int) {

	time.Sleep(time.Millisecond * 100)

	number <- 15 //unblock, jumpto case
}

func channelMessage(message chan string) {

	time.Sleep(time.Millisecond * 100)

	message <- "Learning Go Select" // unblock, jumpto case
}
