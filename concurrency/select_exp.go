package main
import (
  "fmt"
  "time"
)

func main() {

    // create channels
    number := make(chan int)
    message := make(chan string)

    // function call with goroutine

    for i:=0; i < 10; i++ {
        _ = i
        go channelNumber(number)
        go channelMessage(message)
        
    }

  //selects and executes a channel
    // select {
    // case firstChannel := <-number:
    //     fmt.Println("Channel Data:", firstChannel)

    // case secondChannel := <-message:
    //     fmt.Println("Channel Data:", secondChannel)
    // }
    for i:=0; i < 10; i++ {
        _ = i
        loop(number, message)
        
    }

}

func loop(number chan int, message chan string) {
    select {
    case firstChannel := <-number:
      fmt.Println("Channel Data:", firstChannel)

    case secondChannel := <-message:
      fmt.Println("Channel Data:", secondChannel)
    }
}

// goroutine to send data to the channel
func channelNumber(number chan int) {

  // sleeps the process for 2 seconds
  time.Sleep(time.Millisecond*100)

  number <- 15
}

// goroutine to send data to the channel
func channelMessage(message chan string) {

  // sleeps the process by 2 seconds
  time.Sleep(time.Millisecond*100)

  message <- "Learning Go Select"
}