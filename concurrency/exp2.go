package main
 
import (
	"fmt"
	"time"

	"phong/tricks"
)
 
func main() {
	// List of cities.
	cities := []string{"London", "Istanbul", "Berlin", "Madrid"}
 
	// Create unbuffered channel.
	ch := make(chan string)
	
	for index, city := range cities {
		tricks.FormatTwo("loop START", index)

		// Pass channel and city to welcome function and run it as goroutine.
		go welcome(ch, city)
 
		for msg := range ch {
			fmt.Println(msg)
			break
		}

		tricks.FormatTwo("loop END", index)
	}
 
	fmt.Printf("\t\tGOROUTINE main END")
}
 
// welcome accepts write-only channel and city.
func welcome(ch chan<- string, city string) {
	fmt.Printf("\tGOROUTINE write START\n")

	time.Sleep(1 * time.Second)
 
	// Write (send) greeting message to the channel.
	ch <- fmt.Sprintf("read to %s", city)
	fmt.Printf("\tGOROUTINE write END\n")
}