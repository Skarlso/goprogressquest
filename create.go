package main

import (
	"crypto/sha1"
	"fmt"
	"log"
	"net/http"

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
		Race: ELF,
		Cast: MAGE,
		Gold: 0,
		Inventory: Inventory{
			Items: []Item{},
		},
		Level: 1,
		Stats: Stats{
			Strenght:     1,
			Intelligence: 1,
			Luck:         1,
			Perception:   1,
			Agility:      1,
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
