
package main


import(
  "hash/fnv"
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

  h := fnv.New64()
  //Copy is (writer, reader)
  io.Copy(h, f)
  fmt.Println("The sum is: ", h.Sum64())
}
