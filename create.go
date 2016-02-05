package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//create handling the creation of a new character
//curl -H "Content-Type: application/json" -X POST -d '{"name":"asdf"}' http://localhost:8989
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

	charID := c.Param("id")
	var resultCharacter Character
	log.Println("Looking for character with ID:", charID)

	resultCharacter, err := mdb.Load(charID)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{"Error occured while loading character:" + err.Error()})
		return
	}

	c.JSON(http.StatusOK, resultCharacter)
}

func createCharacter(id, name string) Character {
	ch := Character{
		ID:   id,
		Name: name,
		Race: 1,
		Cast: 1,
		Gold: 0,
		Inventory: Inventory{
			Items: []Item{},
		},
		Level: 1,
		//Starts out with everything at 5. Varies randomly later.
		Stats: Stats{
			Strenght:     5,
			Intelligence: 5,
			Luck:         5,
			Perception:   5,
			Agility:      5,
		},
		Body: Body{
			LRing:   Item{},
			RRing:   Item{},
			Armor:   Item{},
			Head:    Item{},
			Weapond: Item{},
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
	races := []Race{}
	if err = json.Unmarshal(dat, &races); err != nil {
		panic(err)
	}

	return races[r.Intn(len(races)-1)].ID
}

//TODO: There is clearly duplication here... To tired though.
func selectRandomCast() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	dat, err := ioutil.ReadFile("casts.json")
	if err != nil {
		panic(err)
	}
	casts := []Cast{}
	if err = json.Unmarshal(dat, &casts); err != nil {
		panic(err)
	}

	return casts[r.Intn(len(casts)-1)].ID
}
