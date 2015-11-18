package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

type DbContext struct {
	Users     *mgo.Collection
	Books  *mgo.Collection
}

func NewDbContext(ip string) DbContext {
	dbSession, err := mgo.Dial(ip)
	if err != nil {
		fmt.Println(err)
	}
	dbContext := DbContext{}
	dbContext.Users = dbSession.DB("IByter").C("Users")
	dbContext.Books = dbSession.DB("IByter").C("Books")
	return dbContext
}
