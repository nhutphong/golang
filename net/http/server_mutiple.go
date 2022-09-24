package main

import (
    "fmt"
    "net/http"
    "sync"
)

func createServer(port int) *http.Server {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "home page at %d", port)
    })

    mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "about page at %d", port)
    })


    server := http.Server{
        Addr:    fmt.Sprintf(":%d", port),
        Handler: mux,
    }

    return &server
}

func main() {
    wg := new(sync.WaitGroup)

    wg.Add(3)

    go func() {
        server := createServer(8081)
        fmt.Println(server.ListenAndServe())
        wg.Done()
    }()

    go func() {
        server := createServer(8082)
        fmt.Println(server.ListenAndServe())
        wg.Done()
    }()

    go func() {
        server := createServer(8083)
        fmt.Println(server.ListenAndServe())
        wg.Done()
    }()

    wg.Wait()
}