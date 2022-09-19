package main

import (
	"fmt"
	"math/rand"
	"time"
)


func generator(msg string) <-chan string {
	ch := make(chan string)
	go func() { 
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}

	}()
	return ch
}

// <-chan string only get the receive value
// fanIn spawns 2 goroutines to reads the value from 2 channels
// then it sends to value to result channel( `c` channel)
func fanIn(c1, c2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for { // infinite loop to read value from channel.
			v1 := <-c1 // read value from c2. This line will wait when receiving value.
			c <- v1
		}
	}()
	go func() {
		for {
			c <- <-c2 // read value from c2 and send it to c
		}
	}()
	return c
}

func fanInSimple(cs ...<-chan string) <-chan string {
	c := make(chan string)
	for _, ci := range cs { // spawn channel based on the number of input channel

		go func(cv <-chan string) { // cv is a channel value
			for {
				c <- <-cv
			}
		}(ci) // send each channel to

	}
	return c
}

func main() {
	// merge 2 channels into 1 channel
	// c := fanIn(generator("Joe"), generator("Ahn"))
	c := fanInSimple(generator("Joe"), generator("Ahn"))

	for i := 0; i < 5; i++ {
		fmt.Println(<-c) // now we can read from 1 channel
	}
	fmt.Println("You're both generator. I'm leaving")
}
