package db

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"os"
)

var (
	Session *mgo.Session
	Mongo   *mgo.DialInfo

	LOGIN_DB    = os.Getenv("LOGIN_DB")
	PASSWORD_DB = os.Getenv("PASSWORD_DB")
	IP          = os.Getenv("IP")
	DB_TYPE     = os.Getenv("DB_TYPE")
)

// Connect connects to mongodb
func Connect() {
	var uri string
	if DB_TYPE == "local" {
		uri = ("mongodb://0.0.0.0:27017/admin")
	} else if DB_TYPE == "server" {
		uri = ("mongodb://" + LOGIN_DB + ":" + PASSWORD_DB + "@" + IP + ":27017/" + "admin")
	}
	mongo, err := mgo.ParseURL(uri)
	s, err := mgo.Dial(uri)
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		panic(err.Error())
	}
	s.SetSafe(&mgo.Safe{})
	fmt.Println("Connected to", uri)
	Session = s
	Mongo = mongo
}
