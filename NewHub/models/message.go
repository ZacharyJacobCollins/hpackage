package models

import "gopkg.in/mgo.v2/bson"

type (
	// User represents the structure of our resource
	Message struct {
		Id     	   bson.ObjectId `json:"id" bson:"_id"`
		Sender		 string     	 `json:"sender" bson:"sender"`
		Username   string        `json:"username" bson:"username"`
		Email  		 string        `json:"email" bson:"email"`
	}
)
