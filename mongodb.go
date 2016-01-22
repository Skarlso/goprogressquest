package main

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Starting mongodb -> mongod --config /usr/local/etc/mongod.conf --fork

//MongoDBConnection Encapsulates a connection to a database
type MongoDBConnection struct {
	session *mgo.Session
}

//Save will save a player using mongodb as a storage medium
func (mdb MongoDBConnection) Save(ch *Character) {
	c := mdb.session.DB("adventure").C("characters")
	err := c.Insert(ch)
	if err != nil {
		panic(err)
	}
}

//Load will load the player using mongodb as a storage medium
func (mdb MongoDBConnection) Load(ID string) (result Character) {
	c := mdb.session.DB("adventure").C("characters")
	err := c.Find(bson.M{"id": ID}).One(&result)
	if err != nil {
		panic(err)
	}
	return
}

//GetSession return a new session if there is no previous one
func (mdb *MongoDBConnection) GetSession() *mgo.Session {
	if mdb.session != nil {
		return mdb.session.Copy()
	}
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session
}
