package main
 
import (
	"fmt"
	// "os"d
	
	"time"
	"phong/tricks"
)
 

 func write(buffer chan int, total int, name string, start int) {
	fmt.Printf("GOROUTINE write() START %s\n", name)
	start++

	for i := 1; i <= total; i++ {
		tricks.FormatTwo("WRITE FOR START", i)
		start *= 10
		buffer <- start
		// i = 4, buffer[2 3] đã đầy nên jumpto goroutine nó connect để READ 
		// i = 1, write=1 lần đầu đã trả cho READ, nên còn lại [2 3]

		tricks.FormatTwo("WRITE FOR END", i)
	}

	// close(buffer)
	fmt.Printf("GOROUTINE write() END %s\n", name)
	fmt.Println()
}

func read(buffer chan int, total int, name string) {
	fmt.Printf("GOROUTINE read() START %s\n", name)

	for i := 1; i <= total; i++ {
		tricks.FormatTwo("READ FOR START", i)
		fmt.Println("index",i, "read", <-buffer)
		tricks.FormatTwo("READ FOR END", i)
	}

	fmt.Printf("GOROUTINE read() END %s\n", name)
}

func SuperGoroutineControl() {

	buffer := make(chan int)  // STEP 2
	defer close(buffer)
	defer fmt.Println("SuperGoroutineControl() END")
	defer time.Sleep(time.Second)


	go write(buffer, 3, "ONE", 7)
	go write(buffer, 3, "TWO", 12)	
	go write(buffer, 3, "THREE", 123)
	// go write(buffer, 3, "FOUR", 11)
	// go write(buffer, 3, "FIVE", 17)

	fmt.Println(cap(buffer))
	go read(buffer, 9, "ONE") //; đủ điều kiện 6 goroutine run xong; sau đó toi SuperGoruotineControl()



	// time.Sleep(time.Second) // phai co de jumpto goroutine CHILD
	// fmt.Println("SuperGoroutineControl() END")
	// close(buffer)
}


func main() {
	
	// SuperGoroutineControl()


	buffer := make(chan int, 12) // STEP 2
	defer close(buffer)
	defer tricks.Format("GOROUTINE MAIN() END")
	defer time.Sleep(time.Second)

	go write(buffer, 3, "ONE", 7)
	go write(buffer, 3, "TWO", 12)	
	go write(buffer, 3, "THREE", 123)
	go write(buffer, 3, "THREE", 332)

	for i := 1; i <= cap(buffer); i++ {
		fmt.Println("index", i, "read", <-buffer)
	}

	//NOT
	// for val := range buffer {
	// 	fmt.Println(len(buffer))
	// 	fmt.Println(val)
	// }


	// time.Sleep(time.Second)
	// tricks.Format("GOROUTINE MAIN() END")
}
 
