package main

import (
	"fmt"
	"log"
	// "time"
)
/*
	Buffrered channel có từ 1 khoảng trống trở lên để chứa dữ liệu, 
	không yêu cầu cả 2 goroutines gửi và nhận cùng phải sẵn sàng cùng lúc. 
	Luồng xử lý chỉ bị block nếu: goroutines write bị block nếu không còn khoảng trống trong channel; 
	goroutines read bị block nếu trong channel không có dữ liệu.

	NOTE buffered channel : có thể write , read cùng 1 goroutine
	or goroutine Parent write<- call thêm goroutine Child anonymous <-read



func main() {
	fmt.Println()
	// unbuffer := make(chan int, 0)
	buffer := make(chan int, 2)

	buffer <- 1
	buffer <- 2

	buffer <- 3 // deadlock / chi write 2 lan vao channel

	log.Println(<-buffer)
	log.Println(<-buffer)
	

	log.Println("end main()")
}
	


*/


/*
	
*/
func goroutine(buffer chan int) {
	buffer <- 1
	buffer <- 2
	buffer <- 3
	buffer <- 4
	buffer <- 5
	//run goroutine thu 2
	go func() { 
		for i:=0; i<10; i++ {
			fmt.Println(<-buffer)
		}
	}()
}



func main() {
	fmt.Println()
	// unbuffer := make(chan int, 0)
	buffer := make(chan int, 5)

	// go goroutine(buffer)



	// buffer <- 1
	// buffer <- 2
	// buffer <- 3
	// buffer <- 4
	// buffer <- 5
	// go func() {
	// 	for i:=0; i<10; i++ {
	// 		fmt.Println(<-buffer)
	// 	}
	// }()


	// write<-,  <-read trên cùng goroutine main()
	// buffer <- 1
	// buffer <- 2
	// buffer <- 3
	// buffer <- 4
	// buffer <- 5

	// fmt.Println(<-buffer)
	// fmt.Println(<-buffer)
	// fmt.Println(<-buffer)
	// fmt.Println(<-buffer)
	// fmt.Println(<-buffer)



	// run goroutine thu 2
	// go func(){
	// 	log.Println("anonymous goroutine 2 START")
	// 	buffer <- 1
	// 	buffer <- 2
	// 	// log.Println(<-buffer)
	// 	// log.Println(<-buffer)
	// 	log.Println("anonymous goroutine 2 END")
	// }()

	
	// log.Println(<-buffer)
	// log.Println(<-buffer)

	log.Println("end main()")
}