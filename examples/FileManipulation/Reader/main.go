//Golang Cat replacement

package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "os"
)

func main()  {
  //my-cat filename
  //os.Args[0] == my-cat
  //os.Args[1] == filename
  f, err := os.Open(os.Args[1])
  if err != nil {
    log.Fatalln("died", err.Error())
  }
  defer f.Close()
  bytes, err := ioutil.ReadAll(f)
  if err != nil {
    log.Fatalln("dead")
  }
  str:=string(bytes)
  fmt.Println(str)
}
