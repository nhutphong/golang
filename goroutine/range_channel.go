package main
 
import (
	"fmt"
	"time"

	"phong/tricks"
)
 
func main() {
	cities := []string{"London", "Istanbul", "Berlin", "Madrid", "Ha Noi"}
 
	
	tricks.Format("gorotine write NOT time.Sleep")
	channel := make(chan string)
	go writeNoTimeSleep(channel, cities)
	read(channel)

 	tricks.Format("goroutine write co time.Sleep")
	ch := make(chan string)
	go write(ch, cities)
	read(ch)
 

	// Indefinitely wait for the channel to return something.
	// for msg := range ch {
	// 	fmt.Println(msg)
	// }
 
	fmt.Printf("\t\tGOROUTINE main END")
} 


func write(ch chan<- string, cities []string) {
	fmt.Printf("\tGOROUTINE write START\n")

	for index, city := range cities {
		fmt.Println("WRITE index", index+1)
		time.Sleep(time.Second)
 
		// Write (send) greeting message to the channel.
		ch <- fmt.Sprintf("to %s", city)
	}
 

	fmt.Printf("\tGOROUTINE write END\n")
	close(ch)
}

func writeNoTimeSleep(ch chan<- string, cities []string) {
	fmt.Printf("\tGOROUTINE write START\n")

	for index, city := range cities {
		fmt.Println("WRITE index", index+1)
 
		ch <- fmt.Sprintf("to %s", city)
	}
 

	fmt.Printf("\tGOROUTINE write END\n")
	close(ch)
}


func read(ch chan string) {
	fmt.Printf("\tGOROUTINE read START\n")
	for msg := range ch {
		fmt.Println("READ", msg)
	}
	fmt.Printf("\tGOROUTINE read END\n")
}

const NOTE string =`
	range channel: chỉ dùng 1 to 1, tức là 1 goroutine write với 1 func read, trên 1 channel(buffer với unbuffer 
		cho kết quả như nhau)

`