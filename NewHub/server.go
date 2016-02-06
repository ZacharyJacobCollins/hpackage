package main

//TODO debug websockets.  Not being embedded

import (
	//Official pacakges
	"net/http"

	//Third party packages
	"github.com/ZacharyJacobCollins/DevHub/models"
	// "github.com/ZacharyJacobCollins/DevHub/controllers"
	// "github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)
//* REMEMBER pulling models from github repo - imported above

//architecture controller -> hubs -> connections

//Minimum one ws0 socket to function.
func main() {
	//handle assets and allow router to access.
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	//create a controller
	contr := models.NewController()

	//Add 3 hubs for testing, startup webserver and controller
	go contr.Run()

	go contr.Serve()

	// Instantiate a new router
	//  r := httprouter.New()
	//
	// // Get a UserController instance
	// mc := controllers.NewMessageController(getSession())
	//
	// // Create a new user
	// r.POST("/message", mc.AddMessage)
	//
	// // Get a user resource
	// r.GET("/message/:id", mc.FindMessage)
	//
	// // Remove an existing user
	// r.DELETE("/message/:id", mc.DeleteMessage)

	// Fire up the server
	// http.ListenAndServe("localhost:3000", r)
}


// getSession creates a new mongo session and panics if connection error occurs
func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection error, is mongo running?
	if err != nil {
		panic(err)
	}

	// Deliver session
	return s
}


//There is a new socket.  default to last one in slice of hubs for everything though.  Why
//TODO map of socket integers to name of chat channels hubs whatever
//TODO make all functions such as NewController part of the struct
//TODO should initialization/methods in structs relaly contain pointers?
//Ex *controller vs controller - should be read only, http://golangtutorials.blogspot.com/2011/06/methods-on-structs.html
