package main

import (
    "fmt"
    "net/http"
)





func main() {
    // DefaultServeMux()
    // NewServeMux()


    handle := new(handle)
    fmt.Println("Server listenning on port 3000 ...")
    fmt.Println(http.ListenAndServe(":3000", handle))
}

type handle struct{}
// implement ServeHTTP()
func (h *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/" {
        w.Write([]byte("Home page"))
    } else if r.URL.Path == "/about" {
        w.Write([]byte("About page"))
    } else {
        w.Write([]byte("Page not found"))
    }
}




func NewServeMux() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", homeHandle)
    mux.HandleFunc("/about", aboutHandle)

    fmt.Println("Server listenning on port 3000 ...")
    fmt.Println(http.ListenAndServe(":3000", mux)) // NewServeMux
}


func DefaultServeMux() {
    http.HandleFunc("/", homeHandle)
    http.HandleFunc("/about", aboutHandle)

    fmt.Println("Server listenning on port 3000 ...")
    fmt.Println(http.ListenAndServe(":3000", nil)) //default
}


func homeHandle(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Home page"))
}

func aboutHandle(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("About page"))
}





// func homeHandle(w http.ResponseWriter, r *http.Request) {
//     switch r.Method {
//     case "GET":
//         w.Write([]byte("Get method"))
//     case "POST":
//         w.Write([]byte("Post method"))
//     case "PUT":
//         w.Write([]byte("Put method"))
//     case "DELETE":
//         w.Write([]byte("Detele method"))
//     }
// }


// func homeHandle(w http.ResponseWriter, r *http.Request) {
//     if r.URL.Path == "/" {
//         w.Write([]byte("Home page"))
//     } else if r.URL.Path == "/about" {
//         w.Write([]byte("About page"))
//     } else {
//         w.Write([]byte("Page not found"))
//     }
// }