package main

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type FrontStation struct {
	Latitude  string        "Latitude"
	Longitude string        "Longitude"
	Id        bson.ObjectId `json:"Id" bson:"_id"`
}

type FrontAlarm struct {
	EventId   bson.ObjectId `json:"EventId" bson:"_id"`
	StCode    bson.ObjectId `json:"StCode" bson:"StCode"`
	DaId      int           `bson:"DaId"`
	DrId      int           `bson:“DrId"`
	AlarmType int           `bson:"AlarmType"`
	OccurTim  time.Time     `bson:"OccurTime"`
	ImageNum  int           `bson:"ImageNum"`
	VideoNum  int           `bson:"VideoNum"`
	Mark      string        `bson:"Mark"`
}

type EventAlarm struct {
	EaId          bson.ObjectId `bson:"_id"`
	EventId       bson.ObjectId `bson:"ObjectId"`
	DbiId         int           `bson:"DbiId"`
	AlarmType     int           `bson:"AlarmTye"`
	AlarmLevel    int           `bson:"AlarmLevel"`
	EaTime        time.Time     `bson:"EaTime"`
	EaValue       string        `bson:"EaValue"`
	EaDescription string        `bson:"EaDescription"`
	Mark          string        `bson:"Mark"`
}
