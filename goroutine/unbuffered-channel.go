package main

import (
	"fmt"
	// "log"
	"time"
	
	"phong/tricks"
)

/*
	SUPER GOROUTINE MAIN() OR PARENT GOROUTINE nên code: defer time.Sleep(time.Second) ở đầu, để chờ CHILD
	GOROUTINE

	các WRITE GOROUTINE nên code: defer close(buffer) or defer close(unbuffer) ở đầu, để tránh
	READ GOROUTINE read buffer empty sẽ gây ra error DEADLOCK 
*/

/*
	chu kỳ chạy của 1 unbuffered-channel: write vs read, or read vs write, tức la có 2 bước:
	vd: gọi là START(write or read) + END(write or read) là 1 chu kỳ 
	khi START sẽ jumpto to END, hết 1 chu kỳ code chạy tiếp, nếu gặp START thì JUMPTO to END 

	tức la gặp SỐ LẺ(START 1 KỲ) THÌ JUMPTO = NHẢY tới goroutine chứa SỐ CHẴN(END 1 CHU KỲ) TIẾP THEO
	gặp SỐ CHẴN(END 1 CHU KỲ) THÌ CỨ RUN QUA NÓ, NẾU GẶP LẠI SỐ LẼ THÌ JUMPTO = NHẢY TIẾP



	block: chờ case or write quá số lần unbuffered sẽ deadlock 
	  select { case }
	  unbuffered channel(write || read)

	Unbuffered channel không có khoảng trống để chứa dữ liệu,
	NOTE: yêu cầu cả 2 goroutines gửi và nhận đều sẵn sàng cùng lúc.

	Khi 1 goroutine gửi dữ liệu vào channel, 
	luồng xử lý sẽ bị block lại cho tới khi có 1 goroutine đọc từ channel này ra.



	unbuffered không có HỒ CHỨA như buffered, nên cho nó xuống CỐNG = (deadlock) 
	NOTE: TIPS AND TRICKS
	STEP 1. goroutine main() is super goroutine: 1 hòn đảo lớn
	STEP 2. unbuffer: là 1 PIPE = 1 đường ống, or hiều là bridge= cây cầu cung được , đẻ nối 2 đảo 
	STEP 3. goroutine ChildLand(): 1 hòn đảo mới mua thêm


	STEP 4. write and read: hiểu là đổ nước vào PIPE và lấy nước ra từ PIPE

	5. LET` GO!!!

	GO:
	1 to 2 to 3 rồi tới 4  //OK 
	tại hòn đảo lơn, nối PIPE tới hòn đảo mới mua, sau đó trao đổi nước giữa 2 đảo

	nếu: 1 to 2 to 4:  //DEADLOCK 
	tại hòn đảo lơn có đường ống=PIPE, khi đổ nước vào và lấy nước ra, là đang trao đổi với CỐNG
	= DEADLOCK, vì PIPE có nối với đảo nào đâu (chưa mua đảo)



	super goroutine main() và Child() thằng nào write or read đều được
	tốt nhất super nên la <-READ    receive=nhận 
			 CHILD nên la WRITE<-   send=gửi



	


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
	defer close(unbuffer)

	unbuffer <- 1 
	unbuffer <- 2
	unbuffer <- 3
	unbuffer <- 4
	// unbuffer <- 5
	// unbuffer <- 6
	// unbuffer <- 7

	//run goroutine thu 2
	// go func() { 
	// 	for i:=0; i<10; i++ {
	// 		fmt.Println(<-unbuffer)
	// 	}
	// }()
}

func goWrite(unbuffer chan int) {
	defer close(unbuffer)

	fmt.Println("write start 1")
	unbuffer <- 1
	fmt.Println("write end 1")
	fmt.Println("write start 2")
	unbuffer <- 2 // jumto 3
	fmt.Println("write end 2")
	fmt.Println("write start 3")
	unbuffer <- 3
	fmt.Println("write end 3")
	fmt.Println("write start 4")
	unbuffer <- 4 // jumto 7
	fmt.Println("write end 4")
	fmt.Println("write start 5")
	unbuffer <- 5
	fmt.Println("write end 5")

	// 
	// close(unbuffer)
}

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


func readUnBufferUseRange(unbuffer chan int) {
	fmt.Println("GOROUTINE readUnBufferUseRange START")

	for item := range unbuffer {
		fmt.Println("read value", item, " for item := range unbuffer")
	}

	fmt.Println("GOROUTINE readUnBufferUseRange END")
}

func f(c chan int) {
    for {
        c <- time.Now().Second()
        time.Sleep(time.Second)
    }
}

func SuperGoroutineControl() {
	/*
		unbuffered channel: code theo pattern dưới; 5 goroutine write với 1 goroutine read
		read = write=15 (trên channel) ; mới đảm bảo 6 goroutine sẽ run xong, sau đó
		mới tới SuperGoroutineControl()
		read > write: vd read=20, thì tràn ra ngoài main() read tới 20 


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

	go read(unbuffer, 30, "ONE")


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


	// go goDeadlock(unbuffer) // write=4; có close(unbuffer)
	// fmt.Println(<-unbuffer)
	// fmt.Println(<-unbuffer)
	// fmt.Println(<-unbuffer)
	// fmt.Println(<-unbuffer)
	// fmt.Println(<-unbuffer) // read=5; NOT DEADLOCK
	// fmt.Println(<-unbuffer) // read=6; 
	// fmt.Println(<-unbuffer) // read=7; 


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