package main

import (
    "fmt"
    "golang.org/x/net/websocket"
    "net/http"
)

func handler(c *websocket.Conn) {
    var s string
    fmt.Fscan(c, &s)
    fmt.Println("Received:", s)
    fmt.Fprint(c, "How do you do?")
}

func main() {
    http.Handle("/", websocket.Handler(handler))
    http.ListenAndServe("localhost:4000", nil)
}
