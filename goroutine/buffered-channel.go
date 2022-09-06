package main

import (
	"fmt"
	"time"
	// "log"
	// "runtime"

	"phong/tricks"
)

/*
	SUPER GOROUTINE MAIN() OR PARENT GOROUTINE nên code: defer time.Sleep(time.Second) ở đầu, để chờ CHILD
	GOROUTINE

	các WRITE GOROUTINE nên code: defer close(buffer) or defer close(unbuffer) ở đầu, để tránh
	READ GOROUTINE read buffer empty sẽ gây ra error DEADLOCK 
*/

/*
	
	goroutines READ bị block(JUMPTO) nếu trong channel không có dữ liệu.
	goroutines WRITE bị block(JUMPTO) nếu không còn khoảng trống trong channel;
	vd: buffer có sức chứa là 2
		nếu chúng ta READ trước thì sẽ JUMPTO to goroutine chứa WRITE để run
		trong goroutine sẽ WRITE liên tục 3 lần =FULL, vì lần đâu tiên trả cho read
		còn lại 2=full, nếu WRITE thêm nữa sẽ jumpto goroutine read


	non-block: code run tiếp không dừng 
	  go myFunc
	  buffered channel(write || read)


	write <-:
		channel <- "write data vào channel"
	<- read:
		content := <-channel

	Buffrered channel có từ 1 khoảng trống trở lên để chứa dữ liệu, 
	không yêu cầu cả 2 goroutines gửi và nhận cùng phải sẵn sàng cùng lúc. 
	


	NOTE buffered channel : có thể write , read cùng 1 goroutine
	goroutine PARENT nên là <-read ;
	goroutine CHILD thì write-<
*/


/*
	
*/

func writeGoroutine(buffer chan int) {
	defer close(buffer)
	//run goroutine thu 2
	go func() { 
		for i:=0; i<10; i++ {
			fmt.Println(<-buffer)
		}
	}() // () declared va run lun
	buffer <- 1
	buffer <- 2
	buffer <- 3
	buffer <- 4
	buffer <- 5
}


func write(buffer chan int) {
	fmt.Println("FUNC write() START")

	for i := 1; i < 5; i++ {
		tricks.FormatTwo("WRITE FOR START", i)
		buffer <- i
		// i = 4, buffer[2 3] đã đầy nên jumpto goroutine nó connect để READ 
		// i = 1, write=1 lần đầu đã trả cho READ, nên còn lại [2 3]

		tricks.FormatTwo("WRITE FOR END", i)
	}

	close(buffer)
	fmt.Println("FUNC write END")
	fmt.Println()
}

func read(buffer chan int) {
	fmt.Println("FUNC read START")

	for i := 1; i < 5; i++ {
		tricks.FormatTwo("READ FOR START", i)
		fmt.Println("read value", <-buffer, "from ch")
		tricks.FormatTwo("READ FOR END", i)
	}

	fmt.Println("FUNC read END")
}

func readBufferUseRange(buffer chan int) {
	fmt.Println("FUNC readBufferUseRange START")

	for item := range buffer {
		fmt.Println("read value", item, " for _ := range buffer")
	}

	fmt.Println("FUNC readBufferUseRange END")

}


// STEP 1
func main() {
	// defer time.Sleep(time.Second)

	tricks.Format("buffered-channel")
	tricks.Format("START GOROUTINE MAIN()")

	// unbuffer := make(chan int, 0)
	buffer := make(chan int, 2) // STEP 2

	// go writeGoroutine(buffer)


	go write(buffer)
	go read(buffer)
	// go readBufferUseRange(buffer)



	// buffer <- 1 
	// buffer <- 2
	// buffer <- 3
	// buffer <- 4
	// buffer <- 5
	// go func() {
	// 	for i:=1; i<10; i++ {
	// 		tricks.FormatTwo("FOR READ START", i)
	// 		fmt.Println(<-buffer) // JUMPTO GOROUTINE MAIN() nen khong DEADLOCK
	// 		tricks.FormatTwo("FOR READ END", i)
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

	
	//
	time.Sleep(time.Second)
	tricks.Format("END GOROUTINE MAIN()")
}