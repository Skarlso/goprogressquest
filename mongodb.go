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

	player := bson.M{"id": ch.ID}
	//TODO: I need to find a better way of doing this. Put the fields into a map[string]interface{}?
	change := bson.M{"$set": bson.M{"stats.strenght": ch.Stats.Strenght, "stats.agility": ch.Stats.Agility,
		"stats.intelligence": ch.Stats.Intelligence, "stats.perception": ch.Stats.Perception,
		"stats.luck": ch.Stats.Luck, "stats.constitution": ch.Stats.Constitution, "hp": ch.Hp, "maxhp": ch.MaxHp,
		"currentxp": ch.CurrentXp, "nextlevelxp": ch.NextLevelXp, "gold": ch.Gold, "level": ch.Level}}
	// log.Println("Update Doc:", string(data))
	err := c.Update(player, change)
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
