package characters

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Starting mongodb -> mongod --config /usr/local/etc/mongod.conf --fork

// MongoDBConnection Encapsulates a connection to a database
type MongoDBConnection struct {
	session *mgo.Session
}

// DB Is a global Db session
var DB Storage

func init() {
	DB = MongoDBConnection{}
}

// // Mdb is a global storage instance
// // TODO: Refactor and remove this
// var Mdb Storage

// Save will save a player using mongodb as a storage medium
func (mdb MongoDBConnection) Save(ch Character) error {
	mdb.session = mdb.GetSession()
	defer mdb.session.Close()
	if _, err := mdb.Load(ch.Name); err == nil {
		return fmt.Errorf("character already exists")
	}
	c := mdb.session.DB("adventure").C("characters")
	err := c.Insert(ch)
	return err
}

// Load will load the player using mongodb as a storage medium
func (mdb MongoDBConnection) Load(Name string) (result Character, err error) {
	mdb.session = mdb.GetSession()
	defer mdb.session.Close()
	c := mdb.session.DB("adventure").C("characters")
	err = c.Find(bson.M{"name": Name}).One(&result)
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
		"currentxp": ch.CurrentXp, "nextlevelxp": ch.NextLevelXp, "gold": ch.Gold, "level": ch.Level,
		"inventory.items": ch.Inventory.Items, "inventory.capacity": ch.Inventory.Capacity,
		"body.armor": ch.Body.Armor, "body.head": ch.Body.Head, "body.lring": ch.Body.LRing,
		"body.rring": ch.Body.RRing, "body.shield": ch.Body.Shield, "body.weapond": ch.Body.Weapond}}
	// log.Println("Update Doc:", string(data))
	err := c.Update(player, change)
	return err
}

// GetSession return a new session if there is no previous one
func (mdb *MongoDBConnection) GetSession() *mgo.Session {
	if mdb.session != nil {
		return mdb.session.Copy()
	}
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal("could not get session to db: ", err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session
}
