package models

import "time"

type User struct {
	Id        string    `json:"id" bson:"id" fake:"{number:1,10}"`
	FirstName string    `json:"first_name" bson:"first_name" fake:"{firstname}"`
	LastName  string    `json:"last_name" bson:"last_name" fake:"{lastname}"`
	Phone     string    `json:"phone" bson:"phone" fake:"{phone}"`
	CreatedAt time.Time `json:"created_at" bson:"created_at" fake:"{year}-{month}-{day}"`
}
