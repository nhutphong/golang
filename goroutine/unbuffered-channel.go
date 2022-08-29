package main

import (
	"fmt"
	"log"
	"phong/tricks"
)
/*

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



	chu kỳ chạy của 1 channel: write read, or read write, tức la có 2 bước; giả xử gọi nó 1 và 2
	nếu có nhiều quá trình writes reads từ channel; nó sẽ run như sau
	bắt đầu run gặp 1 jumto goroutine chứa 2 run qua 2, chạy tiếp chừng nào gặp 3 jumpto tiếp

	tức la gặp SỐ LẺ THÌ JUMPTO = NHẢY tới goroutine chứa SỐ CHẴN TIẾP THEO
	gặp SỐ CHẴN THÌ CỨ RUN QUA NÓ, RUN TIẾP TỤC, NẾU GẶP LẠI SỐ LẼ THÌ JUMPTO = NHẢY TIẾP




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

func goWrite(unbuffer chan int) {

	fmt.Println("write start 1")
	unbuffer <- 1
	fmt.Println("write end 1")
	fmt.Println("write start 2")
	unbuffer <- 2 // jumto
	fmt.Println("write end 2")
	fmt.Println("write start 3")
	unbuffer <- 3
	fmt.Println("write end 3")
	fmt.Println("write start 4")
	unbuffer <- 4 // jumto
	fmt.Println("write end 4")
	fmt.Println("write start 5")
	unbuffer <- 5
	fmt.Println("write end 5")

	// 
	close(unbuffer)
}

//STEP 1
func main() {
	//hien tai co goroutine main
	fmt.Println()
	unbuffer := make(chan int) // STEP 2



	go goWrite(unbuffer) // STEP 3
	// STEP 4
	for i:=1; i<=5; i++ {
		tricks.FormatTwo("FOR READ START", i)
		fmt.Println("for: ", <-unbuffer)
		tricks.FormatTwo("FOR READ END", i)
	}






	// goroutine Child đã sẵn sàng
	// STEP 3
	// go func() { 
	// 	for i:=1; i<=10; i++ {
	// 		tricks.FormatTwo("FOR READ START", i)
	// 		fmt.Println(<-unbuffer)
	// 		tricks.FormatTwo("FOR READ END", i)
	// 	}
	// }()

	// // STEP 4
	// fmt.Println("write start 1")
	// unbuffer <- 1
	// fmt.Println("write end 1")
	// fmt.Println("write start 2")
	// unbuffer <- 2
	// fmt.Println("write end 2")
	// fmt.Println("write start 3")
	// unbuffer <- 3
	// fmt.Println("write end 3")
	// fmt.Println("write start 4")
	// unbuffer <- 4
	// fmt.Println("write end 4")
	// fmt.Println("write start 5")
	// unbuffer <- 5
	// fmt.Println("write end 5")




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