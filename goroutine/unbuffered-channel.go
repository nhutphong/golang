package main

import (
	"fmt"
	"log"
)
/*

	Unbuffered channel không có khoảng trống để chứa dữ liệu,
	yêu cầu cả 2 goroutines gửi và nhận đều sẵn sàng cùng lúc.
	Khi 1 goroutine gửi dữ liệu vào channel, 
	luồng xử lý sẽ bị block lại cho tới khi có 1 goroutine đọc từ channel này ra.


	NOTE: TIPS AND TRICKS
	goroutine main() : 
*/


func goroutine(unbuffer chan int) {
	//run goroutine thu 2
	go func() { 
		for i:=0; i<10; i++ {
			fmt.Println(<-unbuffer)
		}
	}()

	unbuffer <- 1
	unbuffer <- 2
	unbuffer <- 3
	unbuffer <- 4
	unbuffer <- 5
}

func goDeadlock(unbuffer chan int) {
	unbuffer <- 1
	unbuffer <- 2
	unbuffer <- 3
	unbuffer <- 4
	unbuffer <- 5
	close(unbuffer)

	//run goroutine thu 2
	go func() { 
		for i:=0; i<10; i++ {
			fmt.Println(<-unbuffer)
		}
	}()
}



func main() {
	//hien tai co goroutine main
	fmt.Println()
	unbuffer := make(chan int)

	// go goroutine(unbuffer)
	go goDeadlock(unbuffer)
	for i:=0; i<10; i++ {
		fmt.Println(<-unbuffer)
	}



	// 
	// go func() { 
	// 	for i:=0; i<10; i++ {
	// 		fmt.Println(<-unbuffer)
	// 	}
	// }()

	// unbuffer <- 1
	// unbuffer <- 2
	// unbuffer <- 3
	// unbuffer <- 4
	// unbuffer <- 5


	// run goroutine thu 2
	// go func(){
	// 	log.Println("anonymous goroutine 2 START")
	// 	unbuffer <- "Có làm thì mới có ăn" // ok 
	// 	log.Println("anonymous goroutine 2 END")
	// }()

	// value, ok := <-unbuffer
	// log.Println("value: ", value, "ok: ", ok)





	log.Println("end main()")
}