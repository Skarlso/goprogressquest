package main

import (
	// "crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Races is an acumulater for races defined in races.json.
type Races struct {
	Races []Race `json:"races"`
}

// Casts is an acumulater for casts in casts.json.
type Casts struct {
	Casts []Cast `json:"casts"`
}

// create handling the creation of a new character
// curl -H "Content-Type: application/json" -X POST -d '{"name":"asdf"}' http://localhost:8989
func create(c *gin.Context) {

	var newName struct {
		Name string `json:"name"`
	}
	ch := NewCharacter{}

	if err := c.BindJSON(&newName); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{"error while binding newName:" + err.Error()})
		return
	}

	checkSum := sha1.Sum([]byte(newName.Name))
	ch.CharacterID = fmt.Sprintf("%x", checkSum)

	char := createCharacter(ch.CharacterID, newName.Name)

	log.Println("Saving character:", char)
	err := mdb.Save(char)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{"error while saving character:" + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, char)
}

func loadCharacter(c *gin.Context) {

	charName := c.Param("name")
	var resultCharacter Character
	log.Println("Looking for character with ID:", charName)

	resultCharacter, err := mdb.Load(charName)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{"Error occured while loading character:" + err.Error()})
		return
	}

	c.JSON(http.StatusOK, resultCharacter)
}

func createCharacter(id, name string) Character {
	w := Item{Name: "Basic Sword", ID: 1, Armor: 0, Dmg: 10, Weight: 1, Value: 10}
	ch := Character{
		ID:   id,
		Name: name,
		Race: selectRandomRace(),
		Cast: selectRandomCast(),
		Gold: 0,
		Inventory: Inventory{
			Items:    []Item{},
			Capacity: 80,
		},
		Level:       1,
		Hp:          100,
		MaxHp:       100,
		CurrentXp:   0,
		NextLevelXp: 1000,
		//Starts out with everything at 5. Varies randomly later.
		Stats: Stats{
			Strenght:     5,
			Intelligence: 5,
			Luck:         5,
			Perception:   5,
			Agility:      5,
			Constitution: 5,
		},
		Body: Body{
			LRing:   Item{},
			RRing:   Item{},
			Armor:   Item{},
			Head:    Item{},
			Weapond: w,
			Shield:  Item{},
		},
	}

	return ch
}

func selectRandomRace() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	dat, err := ioutil.ReadFile("races.json")
	if err != nil {
		panic(err)
	}
	races := Races{}
	if err = json.Unmarshal(dat, &races); err != nil {
		panic(err)
	}

	return races.Races[r.Intn(len(races.Races)-1)].ID
}

//TODO: There is clearly duplication here...
func selectRandomCast() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	dat, err := ioutil.ReadFile("casts.json")
	if err != nil {
		panic(err)
	}
	casts := Casts{}
	if err = json.Unmarshal(dat, &casts); err != nil {
		panic(err)
	}

	return casts.Casts[r.Intn(len(casts.Casts)-1)].ID
}
