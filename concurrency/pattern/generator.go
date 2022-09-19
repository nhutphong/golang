package main

import (
	"fmt"
	"math/rand"
	"time"
)


func boring(msg string) <-chan string {
	ch := make(chan string)
	
	go func() {
		
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
		close(ch)

	}()
	return ch
}

func main() {

	joe := boring("Joe")
	ahn := boring("Ahn")

	for i := 0; i < 15; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ahn)
	}

	// or we can simply use the for rangch
	// for msg := range joe {
	// 	fmt.Println(msg)
	// }
	fmt.Println("You're both boring. I'm leaving")

}
