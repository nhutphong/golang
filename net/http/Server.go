package main

import (
	"fmt"
	"time"
	"net/http"
)

func main() {
    http.HandleFunc("/", liteHandler)
    server := &http.Server{
        Addr: ":8080",
        ReadTimeout: 10 * time.Second,
        WriteTimeout: 10 * time.Second,
        MaxHeaderBytes: 1 << 20,
    }
    server.ListenAndServe()
}

func liteHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Chào mừng đến lập trình Go cho web!")
}


const NOTE string = `

type Server struct {
    Addr string                             // <Địa chỉ IP/tên miền>:<cổng>
    Handler Handler                    // Đối tượng xử lý yêu cầu
    ReadTimeout time.Duration  // Thời gian tối đa đọc yêu cầu
    WriteTimeout time.Duration // Thời gian tối đa ghi phản hồi
    MaxHeaderBytes int             // Chiều dàu phần đầu tính theo byte, kể cả dòng yêu cầu
    TLSConfig *tls.Config         // Cấu hình TLS, tùy chọn
    TLSNextProto map[string]func(*Server, *tls.Conn, Handler)
    ConnState func(net.Conn, ConnState)
    ErrorLog *log.Logger
}

 `