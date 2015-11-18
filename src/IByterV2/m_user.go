package main

type User struct {
	UserId   int64 `json:"UserId" bson:"_id"`
	Nickname string
    Username string
	Passwd string
    Header string
	IsOnline bool
}
