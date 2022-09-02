package main

import (
	"fmt"
	"time"

	"phong/tricks"
)

func write(buffer chan int) {
	// defer close(buffer)

	for i := 1; i < 3; i++ {
		tricks.FormatTwo("WRITE FOR START", i)
		buffer <- i
		// i = 4, buffer[2 3] đã đầy nên jumpto goroutine nó connect để READ 
		// i = 1, write=1 lần đầu đã trả cho READ, nên còn lại [2 3]

		tricks.FormatTwo("WRITE FOR END", i)
	}
	fmt.Println("FUNC write start")
	close(buffer)
	fmt.Println("FUNC write end")
}

func read(buffer chan int) {
	for i := 1; i < 8; i++ {
		tricks.FormatTwo("READ FOR START", i)
		fmt.Println("read value", <-buffer, "from ch")
		tricks.FormatTwo("READ FOR END", i)
	}
	fmt.Println("FUNC READ start")
	fmt.Println("FUNC READ end")
}



func unwrite(unbuffer chan int) {
	defer close(unbuffer)

	for i := 1; i < 5; i++ {
		tricks.FormatTwo("WRITE FOR START", i)
		unbuffer <- i
		tricks.FormatTwo("WRITE FOR END", i)
	}
	fmt.Println("FUNC write start")
	close(unbuffer)
	fmt.Println("FUNC write end")
}

func unread(buffer chan int) {
	for i := 1; i < 5; i++ {
		tricks.FormatTwo("UNREAD FOR START", i)
		fmt.Println("unread value", <-buffer, "from ch")
		tricks.FormatTwo("UNREAD FOR END", i)
	}
	fmt.Println("FUNC UNREAD start")
	fmt.Println("FUNC UNREAD end")
}



func main() {
	// defer time.Sleep(time.Second)


	buffer := make(chan int,2)
	go write(buffer)
	// go read(buffer)

	// for i := 1; i < 5; i++ {
	// 	tricks.FormatTwo("READ FOR START", i)
	// 	fmt.Println("read value", <-buffer, "from ch")
	// 	// i = 1; buffer chưa có data nên jumpto goroutine khác để write 

	// 	tricks.FormatTwo("READ FOR END", i)
	// 	// time.Sleep(2 * time.Second)
	// }

	// tricks.Format("inside main()")

	// fmt.Println("write data")
	// buffer <- 6
	// fmt.Println(<-buffer)



	for v := range buffer {
		fmt.Println("read value", v, " for _ := range buffer")
		// time.Sleep(2 * time.Second)

	}



	// unbuffer := make(chan int)
	// go unwrite(unbuffer)

	// for i := 1; i < 5; i++ {
	// 	tricks.FormatTwo("READ FOR START", i)
	// 	fmt.Println("read value", <-buffer, "from ch")
	// 	tricks.FormatTwo("READ FOR END", i)
	// 	// time.Sleep(2 * time.Second)
	// }



	time.Sleep(time.Second)
	tricks.Format("END MAIN()")
}
