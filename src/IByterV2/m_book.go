package main

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

type Book struct {
	BookId   bson.ObjectId `json:"Id" bson:"_id"`
	Title string
	CreateTime time.Time
	ModifyTime time.Time
    ReadCount uint32
	ReviewCount uint32
	RecommendCount uint32
    Pay uint32
	Summary string
	UserId  bson.ObjectId `json:"UserId" bson:"userId"`
}
