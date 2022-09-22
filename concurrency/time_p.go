package main
 
import (
	"fmt"
	// "os"
	"time"
	// "math/rand"

)
 
const LOOP = 5
func main() {

	result := make(chan int)

	go func() {
		timeout := Timer(1, "xin chao Timer()")
		count := 0
		for {
			select{
			case msg := <-timeout:
				fmt.Println("msg", msg)
				result<- count
			default:
				count++
			}
		}
	}()

	for i := 0; i < LOOP; i++ {
		fmt.Println("loop", i+1, "result", <-result)
		
	}
	fmt.Println("main() end")
}

// = time.After(time.Duration)
func Timer(delay int, msg string) <-chan string {
	fmt.Println("generator() start")
	ch := make(chan string)

	go func() {
		for i := 0; i < LOOP; i++ {
			time.Sleep(time.Second * time.Duration(delay))
			ch<- msg
		}
		close(ch)
	}()

	fmt.Println("generator() end")
	return ch
}