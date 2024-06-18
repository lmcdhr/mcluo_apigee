package main

import (
    "fmt"
    "net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello from mcluo!")
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Wow! you successfully reach the API!")
}

func main() {
    http.HandleFunc("/hello", helloHandler)
    http.HandleFunc("/api", apiHandler)
    http.ListenAndServe(":8080", nil)
}
