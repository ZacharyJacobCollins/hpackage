package main

import (
  "os"
  "log"
)

func main() {
  f, err := os.Create(os.Args[1])
  if err != nil {
    log.Fatalln("done", err.Error())
  }
  defer f.Close();

  str:= "txt"
  bs := []byte(str)

  _, err = f.Write(bs)
  if err != nil {
    log.Fatalln("done", err.Error())
  }

}
