package main

import (
	"fmt"
	"math/rand"
	"time"
)


func boring(msg string, c chan string) {
	for i := 1; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
	fmt.Println("boring end")
}

func main() {

	c := make(chan string)
	// defer close(c)
	
	go boring("boring!", c)
	go boring("gogo!", c)

	for i := 1; i <= 20; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}

	time.Sleep(time.Second)
	fmt.Println("You're boring. I'm leaving")

}
