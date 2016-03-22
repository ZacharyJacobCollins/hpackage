package chat

import (
	"fmt"
	"html/template"
	"net/http"
	"log"
)

type User struct {
	name 	 string
	password string
	image    string
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("html/login.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		u := User{name: r.Form["username"][0], password: r.Form["password"][0], image:r.Form["image"][0]}
		log.Print(u.name)
	}
}
