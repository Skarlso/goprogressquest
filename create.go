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
	return selectRandomAffiliation("races")

}

func selectRandomCast() int {
	return selectRandomAffiliation("casts")
}

func selectRandomAffiliation(t string) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	dat, err := ioutil.ReadFile(t + ".json")
	if err != nil {
		panic(err)
	}
	var m map[string]interface{}
	if err = json.Unmarshal(dat, &m); err != nil {
		panic(err)
	}

	log.Println(m)
	elements := m[t].([]interface{})
	elem := elements[r.Intn(len(elements)-1)]
	return int(elem.(map[string]interface{})["id"].(float64))
}
