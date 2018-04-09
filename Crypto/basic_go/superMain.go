package main

import (
    "net/http"
    L "../basic_go/lib"
)

func main() {
    http.HandleFunc("/", L.Demo())
    http.ListenAndServe(":8080", nil)
}
