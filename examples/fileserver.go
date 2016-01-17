package main

import (
    "net/http"
    "fmt"
)

func main() {
    fmt.Println("Library is up... ")
    http.Handle("/", http.FileServer(http.Dir("/usr/share/nginx/html/library/files")))
    http.ListenAndServe(":80", nil)
    fmt.Println("Library server failed... ")
}
