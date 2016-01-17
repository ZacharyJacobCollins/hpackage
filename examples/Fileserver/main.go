package main

import (
  "net/http"
  "fmt"
  "log"
  "os"
)

func prompt() (string, string){
  var port string
  var dir string

  //Grab the port, if an empty string make it equal to default :0080
  fmt.Println("Enter a port to run the Fileserver (Press enter for default port 80) ...")
  fmt.Scanln(&port)
    if port=="" {
      port = ":http"
    }
  //Grab the current directory, if it's empty use current working directory
  fmt.Println("Enter a directory to display (Press enter to use current directory) ...")
  fmt.Scanln(&dir)
    if dir=="" {
      cwd, err := os.Getwd()
        if err != nil {
          log.Printf("Failure retrieving current directory")
          log.Fatal(err)
        }
      dir = cwd
    }
  return port, dir
}

func main() {
  port, dir := prompt();
  log.Printf("Fileserver is sharing directory %s on %s...\n", "/", port)
  fs := http.FileServer(http.Dir(dir))
  http.Handle("/", fs)
  http.ListenAndServe(port, fs)
}

//TODO validation for correct entered directory.
