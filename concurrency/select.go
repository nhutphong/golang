package main

import (
	"fmt"
	"phong/tricks"
	"time"
)

func action(ch chan int, quit chan bool) {
	tricks.Format("action() start")

	for i := 0; i < 10; i++ {
		fmt.Println("go: ", <-ch)
	}

	fmt.Println("action() end")
	quit <- true

}

func fibonacci(ch chan int, quit chan bool, result chan int) {
	x, y, count := 0, 1, 0
	fmt.Println("\tfibonacci() start")

	dem := 1
	for {
		fmt.Printf("\t\tfor fib start%v\n", dem)

		select {
		case ch <- y:
			x, y = y, x+y
			count++
			fmt.Println("\t\t\tcase write ch<- count: ", count)

		case <-quit:
			fmt.Println("\tquit")
			fmt.Println("\tfibonacci() end")
			result <- y
			return //break for
		}

		fmt.Printf("\t\tfor fib end%v\n", dem)
		dem++
	}

}

const (
	LOOP int = 5
)

func main() {
	ch := make(chan int, 1)
	quit := make(chan bool)
	result := make(chan int)
	go func() {
		fmt.Println("\tMAIN READ START")

		for i := 1; i < LOOP; i++ {
			fmt.Printf("\t\tMAIN() for start%v\n", i)

			fmt.Println("\t\tMAIN() go ch =  ", <-ch)

			fmt.Printf("\t\tMAIN() for end%v\n", i)
			fmt.Println()
		}

		fmt.Println("\tMAIN READ END")
		quit <- true
	}()

	// go action(ch, quit)
	go fibonacci(ch, quit, result)

	fmt.Println("main() middle")
	fmt.Println("result", <-result)
	time.Sleep(time.Second)
	tricks.Format("main() end")
}
