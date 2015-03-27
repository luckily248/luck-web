package models

import (
	"gopkg.in/mgo.v2"
	"log"
)

type BaseDBmodel struct {
	session *mgo.Session
	db      *mgo.Database
	c       *mgo.Collection
}

func (this *BaseDBmodel) DBname() string {
	return "filesys"
}

func (this *BaseDBmodel) init() {
	newsession, err := mgo.Dial("")
	if err != nil {
		log.Fatal("mgo init error")
		panic(err)
	}
	this.session = newsession
	this.session.SetMode(mgo.Monotonic, true)
	this.db = this.session.DB(this.DBname())
}
