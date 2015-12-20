package main

import (
  "os"
  "log"
  "fmt"
)

func main()  {
  f, err := os.Open("hello.txt")
  if err != nil {
    log.Fatalln("died")
  }
  defer f.Close()

  bytes, err := ioutill.ReadAll(f)

  if err != nil {
    log.Fatalln("dead")
  }

  fmt.Println(bytes)
}
