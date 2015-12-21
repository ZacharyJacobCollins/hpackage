//Computes md5 hash for a file
package main


import(
  "crypto/md5"
  "os"
  "io"
  "fmt"
)

func main() {
  f, err := os.Open(os.Args[1])
  if err != nil {
    panic(err)
  }
  defer f.Close()

  h := md5.New()
  //Copy is (writer, reader)
  io.Copy(h, f)
  fmt.Printf("The sum is: %x", h.Sum(nil))
}
