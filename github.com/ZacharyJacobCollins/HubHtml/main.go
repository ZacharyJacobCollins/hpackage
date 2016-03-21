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
  http.Handle("/html/", http.StripPrefix("/html/", http.FileServer(http.Dir("html"))))
  http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./html/"))))
	http.ListenAndServe(":1337", nil)
}
