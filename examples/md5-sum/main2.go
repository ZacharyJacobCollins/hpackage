//Computes md5 hash for a directory

package main

import (
  "os"
  "fmt"
  "path/filepath"
)

func main() {
  filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
    fmt.Println(info)
    return nil
  })
/*
  f, err := os.Open(os.Args[1])
  if err != nil {
    panic(err)
  }
  defer f.Close()

  h:= md5.New()
  io.Copy(h, f)
  fmt.Pringf("%x\n", h.Sum(nil))
*/
}
