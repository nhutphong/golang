package main
 
import (
	"fmt"
	"os"
	"time"
)
 
func main() {
	// List of cities.
	cities := []string{"London", "Istanbul", "Berlin", "Madrid"}
 
	// Create unbuffered channel.
	ch := make(chan string)
 
	// Pass channel and cities to welcome function and run it as goroutine.
	go welcome(ch, cities)
 
	// Indefinitely wait for the channel to return something.
	for {
		select {
		// If something was received through the channel, print it.
		case msg := <-ch:
			fmt.Println(msg)
		// If nothing was received through the channel for 5 seconds, timeout.
		case <-time.After(5 * time.Second):
			fmt.Println("timeout")
			os.Exit(0)
		}
	}
}
 
// welcome accepts write-only channel and list of cities.
func welcome(ch chan<- string, cities []string) {
	// Loop through the cities.
	for _, city := range cities {
		// Sleep 2 seconds before writing to the channel below.
		time.Sleep(1 * time.Second)
 
		// Write (send) greeting message to the channel.
		ch <- fmt.Sprintf("Welcome to %s", city)
	}
}