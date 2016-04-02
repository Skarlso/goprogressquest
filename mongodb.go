package main

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Starting mongodb -> mongod --config /usr/local/etc/mongod.conf --fork

// MongoDBConnection Encapsulates a connection to a database
type MongoDBConnection struct {
	session *mgo.Session
}

// Save will save a player using mongodb as a storage medium
func (mdb MongoDBConnection) Save(ch Character) error {
	mdb.session = mdb.GetSession()
	defer mdb.session.Close()
	c := mdb.session.DB("adventure").C("characters")
	err := c.Insert(ch)
	log.Println("Saving character:", ch)
	return err
}

// Load will load the player using mongodb as a storage medium
func (mdb MongoDBConnection) Load(ID string) (result Character, err error) {
	mdb.session = mdb.GetSession()
	defer mdb.session.Close()
	c := mdb.session.DB("adventure").C("characters")
	err = c.Find(bson.M{"id": ID}).One(&result)
	return result, err
}

// Update update player
func (mdb MongoDBConnection) Update(ch Character) error {
	mdb.session = mdb.GetSession()
	defer mdb.session.Close()
	c := mdb.session.DB("adventure").C("characters")

	data, err := bson.Marshal(&ch)
	if err != nil {
		panic(err)
	}

	player := bson.M{"id": ch.ID}
	change := bson.M{"$set": data}
	err = c.Update(player, change)
	log.Println("Updating character:", ch)
	return err
}

// GetSession return a new session if there is no previous one
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
