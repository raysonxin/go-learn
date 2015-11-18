package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

type DbContext struct {
	FrontStations *mgo.Collection
	EventAlarms   *mgo.Collection
	FrontAlarms   *mgo.Collection
}

func NewDbContext(ip string) DbContext {
	dbSession, err := mgo.Dial(ip)
	if err != nil {
		fmt.Println(err)
	}
	dbContext := DbContext{}
	dbContext.FrontStations = dbSession.DB("MonitorSystem").C("FrontStations")
	dbContext.FrontAlarms = dbSession.DB("MonitorSystem").C("FrontAlarms")
	dbContext.FrontAlarms = dbSession.DB("MonitorSystem").C("FrontAlarms")
	return dbContext
}
