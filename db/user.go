package db

import "gopkg.in/mgo.v2/bson"

// User model for user
type User struct {
	ID       bson.ObjectId `json:"id"    bson:"_id,omitempty"`
	Name     string        `json:"name"  bson:"name"`
	Email    string        `json:"email" bson:"email"`
	Password string        `json:"-"     bson:"password"`
	Image    Image         `json:"image" bson:"image"`
}
