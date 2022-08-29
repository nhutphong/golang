package main

import (
    "fmt"
    "time"
)


func one(ch chan string, msg string) {
    ch <-msg
}

// Đẩy dữ liệu cho channel và chờ 4 giây
func data1(ch chan string) {  
    time.Sleep(4 * time.Second)
    ch <- "from data1()"
}
 
// Đẩy dữ liệu cho channel và chờ 2 giây
func data2(ch chan string) {  
    time.Sleep(2 * time.Second)
    ch <- "from data2()"
}


func main() {
    messages := make(chan string)
    signals := make(chan string)

    // select {
    // case msg := <-messages:
    //     fmt.Println("received message", msg)
    // default:
    //     fmt.Println("no message received")
    // }

    // msg := "hi"
    go one(messages, "xin chao golang")
    go one(signals, "xin chao signals")

    select {
    case msg := <-messages:
        fmt.Println("sent message", msg)
    case sig := <-signals:
        fmt.Println("sent message", sig)
    default:
        fmt.Println("no message sent")
    }

    // select {
    // case msg := <-messages:
    //     fmt.Println("received message", msg)
    // case sig := <-signals:
    //     fmt.Println("received signal", sig)
    // default:
    //     fmt.Println("no activity")
    // }




    // Tạo các biến channel để truyền giá trị string
    chan1 := make(chan string)
    chan2 := make(chan string)
    
    // Gọi các goroutine cùng với các biến channel
    go data1(chan1)
    go data2(chan2)
    
    //Cả hai câu lệnh case đều đợi dữ liệu trong chan1 hoặc chan2
    //chan2 nhân dữ liệu trước do nó chỉ chờ 2 giây trong data2().
    //Nên case thứ 2 sẽ chạy và thoát khỏi khối select
    select {
      case x := <-chan1:
        fmt.Println(x)
      case y := <-chan2:
        fmt.Println(y)
      default:
        fmt.Println("tao la default")
    }

}