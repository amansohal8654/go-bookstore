package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id      bson.ObjectId `json:"id" bson: "id"`
	Name    string        `json:"name" bson: "name"`
	Genders string        `json:"genders" bson: "genders"`
	Age     int           `json:"age" bson: "age"`
}
