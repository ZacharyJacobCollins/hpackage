package main

import (
   "net/http"
   "fmt"
)

func HandleHello(w http.ResponseWriter, r *http.Request) {
  fmt.Println("method:", r.Method) //get request method
  if r.Method == "POST" {
    fmt.Print("Posted")
  } else if r.Method == "GET" {
    fmt.Print("GET")
  } else {
    fmt.Print("MISSED")
  }
}

func main() {
  http.HandleFunc("/test", HandleHello)
  http.ListenAndServe(":1337", nil)
}
