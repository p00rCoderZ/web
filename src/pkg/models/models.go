package models

import (
	"errors"
	"time"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
	ErrBadRequest         = errors.New("models: bad request")
)

//Post ...
// type Post struct {
// 	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
// 	Title   string             `json:"title,omitempty" bson:"title,omitempty"`
// 	Content string             `json:"content,omitempty" bson:"content,omitempty"`
// 	Created time.Time          `json:"created,omitempty" bson:"created,omitempty"`
// }
//Post is a struct
type Post struct {
	ID     int    `json:"_id,omitempty" `
	UserId string `json:"userId,omitempty" `
	Type   string `json:"type, omitempty`
	Title  string `json:"title,omitempty" `
	// UserId  string    `json:"userId,omitempty" `
	Content string    `json:"content,omitempty" `
	Tags    []Tag     `json:"tags,omitempty"`
	Status  bool      `json:"status,omitempty"`
	Created time.Time `json:"created,omitempty" `
}

type Tag struct {
	ID int `json:"_id,omitempty" `
}

// type User struct {
// 	ID             primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
// 	Name           string             `json:"name,omitempty" bson:"name,omitempty"`
// 	Surname        string             `json:"surname,omitempty" bson:"surname,omitempty"`
// 	Email          string             `json:"email,omitempty" bson:"email,omitempty"`
// 	HashedPassword []byte             `json:"hashedPassword,omitempty" bson:"hashedPassword,omitempty"`
// 	Created        time.Time          `json:"created,omitempty" bson:"created,omitempty"`
// }

//User is struct
type User struct {
	ID         int    `json:"id,omitempty" `
	Name       string `json:"name,omitempty" `
	Surname    string `json:"surname,omitempty" `
	Nick       string `json:"nick,omitempty"`
	Email      string `json:"email,omitempty" `
	Password   string `json:"password,omitempty"`
	SoftDelete bool   `json:"softDelete, omitempty"`
	// Created        time.Time          `json:"created,omitempty" bson:"created,omitempty"`
}
