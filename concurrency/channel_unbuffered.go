package main

import (
	"fmt"
	// "log"
	"time"
	
	"phong/tricks"
)

const NOTE string = `
	SUPER GOROUTINE MAIN() OR PARENT GOROUTINE nên code: defer time.Sleep(time.Second) ở đầu, để chờ CHILD
	GOROUTINE

	các WRITE GOROUTINE nên code: defer close(buffer) or defer close(unbuffer) ở đầu, để tránh
	READ GOROUTINE read buffer empty sẽ gây ra error DEADLOCK 
*/

/*
	Unbuffered channel sẽ block goroutine khi write channel 
	Lấy dữ liệu từ empty unbuffered channel sẽ block goroutine

	unbuffered := make(chan int)
	unbuffered <- "hello" //send=write chi 1 lan, la BLOCK goroutine current;


	unbuffered := make(chan int)
	<-unbuffered // receive=read empty channel; BLOCK goroutine current



	block: chờ case or write quá số lần unbuffered sẽ deadlock 
	  select { case }
	  unbuffered channel(write || read)

	Unbuffered channel không có khoảng trống để chứa dữ liệu,
	NOTE: yêu cầu cả 2 goroutines gửi và nhận đều sẵn sàng cùng lúc.


`


func write(unbuffer chan int, total int, name string) {
	fmt.Printf("GOROUTINE write() START %s\n", name)

	for i := 1; i <= total; i++ {
		tricks.FormatTwo("WRITE FOR START", i)
		unbuffer <- i
		// i = 4, unbuffer[2 3] đã đầy nên jumpto goroutine nó connect để READ 
		// i = 1, write=1 lần đầu đã trả cho READ, nên còn lại [2 3]

		tricks.FormatTwo("WRITE FOR END", i)
	}

	// close(unbuffer)
	fmt.Printf("GOROUTINE write() END %s\n", name)
	fmt.Println()
}

func read(unbuffer chan int, total int, name string) {
	fmt.Printf("GOROUTINE read() START %s\n", name)

	for i := 1; i <= total; i++ {
		tricks.FormatTwo("READ FOR START", i)
		fmt.Println("read value", <-unbuffer, "from ch")
		tricks.FormatTwo("READ FOR END", i)
	}

	fmt.Printf("GOROUTINE read() END %s\n", name)
}

func f(c chan int) {
    for {
        c <- time.Now().Second()
        time.Sleep(time.Second)
    }
}

func SuperGoroutineControl() {
	/*
		unbuferred channel:read = write bắt buộc
		unbuffered channel: code theo pattern dưới; 5 goroutine write với 1 goroutine read
		read = write=15 (trên channel) bat buoc ; mới đảm bảo 6 goroutine sẽ run xong, sau đó
		mới tới SuperGoroutineControl()
		read < write: code bình thường vẫn được, nhưng check -race sẽ lỗi DATA RACE 
		read > write : vd read=20, thì tràn ra ngoài main() read tới 20 (van con thg GOROUTINE PARENT de overflow)
		read > write: neu main() la noi read thi se DEADLOCK, con goroutine nao nua dau ma overflow


		các SuperGoroutineControl(): nên chứa:
			buffer := make(chan int, 3) // STEP 2
			defer close(buffer)
			defer fmt.Println("SuperGoroutineControl() END")
			defer time.Sleep(time.Second)

		còn main() chỉ:
			defer time.Sleep(time.Second)
	*/

	
	unbuffer := make(chan int) // STEP 2
	// defer close(unbuffer)
	// defer fmt.Println("SuperGoroutineControl() END")
	// defer time.Sleep(time.Second)

	go write(unbuffer, 3, "ONE")
	go write(unbuffer, 3, "TWO")
	go write(unbuffer, 3, "THREE")
	go write(unbuffer, 3, "FOUR")
	go write(unbuffer, 3, "FIVE")

	go read(unbuffer, 15, "ONE")


	// go read(unbuffer, 10, "ONE")
	// go read(unbuffer, 5, "ONE")

	// go read(unbuffer, 5, "ONE")
	// go read(unbuffer, 5, "ONE")
	// go read(unbuffer, 5, "ONE")


	time.Sleep(time.Second) // phai co de jumpto goroutine read || write, ; se end sau cung  
	fmt.Println("SuperGoroutineControl() END")
	close(unbuffer)
}


//STEP 1
func main() {


	//hien tai co goroutine main
	tricks.Format("unbuffered-channel")
	tricks.Format("START GOROUTINE MAIN()")

	// unbuffer := make(chan int) // STEP 2
	// // defer close(unbuffer)
	// // defer time.Sleep(time.Second)

	// go write(unbuffer, 3, "ONE")
	// go write(unbuffer, 3, "TWO")
	// go write(unbuffer, 3, "THREE")
	// go write(unbuffer, 3, "FOUR")
	// go write(unbuffer, 3, "FIVE")

	// go read(unbuffer, 15, "ONE")

	// go read(unbuffer, 5, "TWO")
	// go read(unbuffer, 5, "THREE")

	// go readUnBufferUseRange(unbuffer)

	SuperGoroutineControl()


	// DEADLOCK
	// for item := range unbuffer {
	// 	fmt.Println("read value", item, " for item := range unbuffer")
	// }


	// go goWrite(unbuffer) // STEP 3
	// // STEP 4
	// for i:=1; i<=7; i++ {
	// 	tricks.FormatTwo("FOR READ START", i)
	// 	fmt.Println("for: ", <-unbuffer)
	// 	tricks.FormatTwo("FOR READ END", i)
	// }


	// goroutine Child đã sẵn sàng
	// STEP 3
	// go func() { 
	// 	for i:=1; i<=10; i++ {
	// 		tricks.FormatTwo("FOR READ START", i)
	// 		fmt.Println(<-unbuffer)
	// 		tricks.FormatTwo("FOR READ END", i)
	// 	}
	// }()


	// run goroutine thu 2
	// go func(){
	// 	log.Println("anonymous goroutine 2 START")
	// 	unbuffer <- "Có làm thì mới có ăn" // ok 
	// 	log.Println("anonymous goroutine 2 END")
	// }()

	// value, ok := <-unbuffer
	// log.Println("value: ", value, "ok: ", ok)

	time.Sleep(time.Second)
	tricks.Format("END GOROUTINE MAIN()")
	// close(unbuffer)
}