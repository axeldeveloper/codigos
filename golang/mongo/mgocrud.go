package main

import (
	"github.com/globalsign/mgo"
	"labix.org/v2/mgo/bson"
)

// Now on running mgocrud generate main.go

type User struct {
	ID   bson.ObjectId `bson:"_id,omitempty"`
	Name string        `bson:"name"`
}

func (u *User) Create(mgosession *mgo.Session, database, collection string) error {
	err := mgosession.DB(database).C(collection).Insert(u)
	if err != nil {
		return err
	}
	return err
}
func (u *User) Read(mgosession *mgo.Session, database, collection string, selector bson.M) (*[]User, error) {
	var results []User
	err := mgosession.DB(database).C(collection).Find(selector).All(&results)
	if err != nil {
		return nil, err
	}
	return &results, nil
}
func (u *User) Update(mgosession *mgo.Session, database, collection string, selector, change bson.M) error {
	err := mgosession.DB(database).C(collection).Update(selector, bson.M{"$set": change})
	if err != nil {
		return err
	}
	return nil
}
func (u *User) Delete(mgosession *mgo.Session, database, collection string, selector bson.M) (*mgo.ChangeInfo, error) {
	info, err := mgosession.DB(database).C(collection).RemoveAll(selector)
	if err != nil {
		return nil, err
	}
	return info, nil
}
