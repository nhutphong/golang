package main

import "fmt"

func main() {
	c := make(chan string, 2)

	trySend := func(v string) {
		select {
		case c <- v:
		default: // go here if c is full.
		}
	}

	tryReceive := func() string {
		select {
		case v := <-c: return v
		default: return "-" // go here if c is empty
		}
	}

	trySend("Hello!") // succeed to send
	trySend("Hi!")    // succeed to send
	
	trySend("Bye!")

	
	fmt.Println(tryReceive()) // Hello!
	// fmt.Println(tryReceive()) // Hi!
	
	fmt.Println(tryReceive()) // -
}