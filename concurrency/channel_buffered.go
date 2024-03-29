package main

import (
	"fmt"
	"time"

	// "log"
	// "runtime"

	"phong/tricks"
)

const NOTE string = `
	buffer-channel use khi: chung ta khong bat buoc phai lay het values duoc write vao in channel:
		dau tien cu cho cap-buffer = write
		luc do, read <= write ():
			read chi can 1 lan van ok(thoa man dk, cac child goroutines van end truoc supergoroutine)

		read > write: co the run, nhung ko nen dung

	Buffered Channel sẽ block goroutine hiện tại nếu vượt sức chứa
	Lấy dữ liệu từ empty buffered channel sẽ block goroutine

	buffered := make(chan int,2)
	buffered <- "hello"
	buffered <- "world"
	buffered <- "hello world" //write=3; block goroutine current; vuot qua suc chua


	buffered := make(chan int,2)
	<-buffered // read empty channel; block 

	vd: buffer có sức chứa là 2
		nếu chúng ta READ trước thì sẽ JUMPTO to goroutine chứa WRITE để run
		trong goroutine sẽ WRITE liên tục 3 lần =FULL, vì lần đâu tiên trả cho read
		còn lại 2=full, nếu WRITE thêm nữa sẽ jumpto goroutine read


	non-block: code run tiếp không dừng 
	  go myFunc
	  buffered channel(write || read)


	write<-:
		channel <- "write data vào channel"
	<-read:
		content := <-channel


	NOTE buffered channel : có thể write , read cùng 1 goroutine
	goroutine PARENT nên là <-read ;
	goroutine CHILD thì write-<
`

/*

 */

func writeGoroutine(buffer chan int) {
	defer close(buffer)
	//run goroutine thu 2
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-buffer)
		}
	}() // () declared va run lun
	buffer <- 1
	buffer <- 2
	buffer <- 3
	buffer <- 4
	buffer <- 5
}

func write(buffer chan int, total int, name string) {
	fmt.Printf("GOROUTINE write() START %s\n", name)

	for i := 1; i <= total; i++ {
		tricks.FormatTwo("WRITE FOR START", i)
		buffer <- i
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
		fmt.Println("index", i, "read", <-buffer)
		tricks.FormatTwo("READ FOR END", i)
	}

	fmt.Printf("GOROUTINE read() END %s\n", name)
}

func SuperGoroutineControl() {
	/*
		buffer := make(chan int, 9) // buffer-cap=9
		cho cap(buffer) = write cho dễ, lúc đó muốn read <= write bao nhiêu thì lấy,
		không bị lỗi DATA RACE
		or cap(buffer) + read = write

		buffered channel: code theo pattern dưới; 5 goroutine write với 1 goroutine read
		read+cap(channel) = write=15 (trên channel) ; mới đảm bảo 6 goroutine sẽ run xong, sau đó
		mới tới SuperGoroutineControl()

		note: cu write = buffer-cap la duoc, tu do read <= write deu ok
		read > write, tuyet doi khong nen code nhu vay, se effect (ket qua khong mong muon)

		các SuperGoroutineControl(): nên chứa:
			buffer := make(chan int, 3) // STEP 2
			defer close(buffer)
			defer fmt.Println("SuperGoroutineControl() END")
			defer time.Sleep(time.Second)

		còn main() chỉ:
			defer time.Sleep(time.Second)
	*/

	// cap(channel) >= 5
	buffer := make(chan int, 15) // STEP 2
	defer close(buffer)
	defer fmt.Printf("\tSuperGoroutineControl() END\n")
	defer time.Sleep(time.Second)

	// write buffer = 15 = cap = 15 // đủ điều kiện end 6 goroutines trước SuperGoroutine()
	go write(buffer, 3, "ONE")
	go write(buffer, 3, "TWO")
	go write(buffer, 3, "THREE")
	go write(buffer, 3, "FOUR")
	go write(buffer, 3, "FIVE")

	// chỉ read 10 data
	go read(buffer, 10, "ONE") //; đủ điều kiện 6 goroutine run xong(cap-buffer=15); sau đó toi SuperGoruotineControl()

	// go read(buffer, 10, "ONE")
	// go read(buffer, 2, "ONE")

	// go read(buffer, 3, "ONE")
	// go read(buffer, 4, "ONE")
	// go read(buffer, 5, "ONE")

	//time.Sleep(time.Second) // phai co de jumpto goroutine CHILD
	// fmt.Println("SuperGoroutineControl() END")
	//close(buffer)
}

// STEP 1
func main() {

	tricks.Format("buffered-channel")
	tricks.Format("START GOROUTINE MAIN()")

	SuperGoroutineControl()

	// unbuffer := make(chan int, 0)

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

	//
	time.Sleep(time.Second)
	tricks.Format("GOROUTINE MAIN() END")
	// close(buffer)
}
