package main

import (
    "net/http" //package for http based web programs
    "fmt"
    "log"
)

func handler(w http.ResponseWriter, r *http.Request) {
      var messages []string
      var input string
      fmt.Scanf("%s", &input)
      var messages []string

      messages = append(messages,input)
      for _, r:= range messages{
          fmt.Fprintf(w, "<h1>"+r+"</h1></br>")
          log.Print(r+"\n")
      }
}

func main() {
      http.HandleFunc("/", handler) // redirect all urls to the handler function
      http.ListenAndServe("localhost:1313", nil) // listen for connections at port 9999 on the local machine
}


//Things can only be printed on html pages, like doc.write.  Errases everything else.
//Need to stick in a channel, and write the channel.
