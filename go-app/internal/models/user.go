package models

import "time"

type User struct {
	Id        uint64    `json:"id" bson:"id"`
	FirstName string    `json:"first_name" bson:"first_name"`
	LastName  string    `json:"last_name" bson:"last_name"`
	Phone     string    `json:"phone" bson:"phone"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
}
