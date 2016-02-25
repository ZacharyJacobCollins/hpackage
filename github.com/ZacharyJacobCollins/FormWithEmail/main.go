package main

import (
  "github.com/bmizerany/pat"
  "html/template"
  "log"
  "net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
  render(w, "templates/index.html", nil)
}

func send(w http.ResponseWriter, r *http.Request) {
  msg := &Message{
    Email: r.FormValue("email"),
    Content: r.FormValue("content"),
  }

  if msg.Validate() == false {
    render(w, "templates/index.html", msg)
    return
  }

  if err := msg.Deliver(); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  http.Redirect(w, r, "/confirmation", http.StatusSeeOther)
}

func confirmation(w http.ResponseWriter, r *http.Request) {
  render(w, "templates/confirmation.html", nil)
}

func render(w http.ResponseWriter, filename string, data interface{}) {
  tmpl, err := template.ParseFiles(filename)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
  if err := tmpl.Execute(w, data); err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func main() {
  mux := pat.New()
  mux.Get("/", http.HandlerFunc(index))
  mux.Post("/", http.HandlerFunc(send))
  mux.Get("/confirmation", http.HandlerFunc(confirmation))

  log.Println("Listening...")
  http.ListenAndServe(":3000", mux)
}
