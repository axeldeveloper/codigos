package main

// using packeges
// https://godoc.org/github.com/globalsign/mgo/bson

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

//const MongoDb details
const (
	hosts      = "localhost:27017"
	database   = "movies_db"
	username   = ""
	password   = ""
	collection = "userdetails"
)

type UserDetails struct {
	_id   bson.ObjectId `bson:"_id,omitempty"`
	name  string
	phone string
}

func main() {

	info := &mgo.DialInfo{
		Addrs:    []string{hosts},
		Timeout:  60 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}

	session, err1 := mgo.DialWithInfo(info)
	if err1 != nil {
		panic(err1)
	}

	col := session.DB(database).C(collection)
	datab := session.DB(database)

	count, err2 := col.Count()
	if err2 != nil {
		panic(err2)
	}

	fmt.Println("Database Name:", datab.Name)
	fmt.Println("Collection FullName:", col.FullName)
	fmt.Println(fmt.Sprintf("Documents count: %d", count))

	var userDetail []bson.M
	_ = col.Find(nil).All(&userDetail)
	for _, v := range userDetail {
		fmt.Println(v)
	}

}
