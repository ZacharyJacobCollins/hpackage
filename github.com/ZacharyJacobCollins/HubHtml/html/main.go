package main

import (
  "net/http"

  //Third party packages
  // "github.com/ZacharyJacobCollins/Wiki/wiki"
  "github.com/zacharyjacobcollins/HubHtml/chat"
)

func main() {
  // w := wiki.NewWiki();  w.Run();
  c := chat.NewChat();  c.Run(3);
  http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./html/"))))
  //http.HandleFunc("/login", chat.Login)
  http.ListenAndServe(":1337", nil)
}
