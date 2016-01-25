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
	log.Printf("Created character sha hash: %v", ch.CharacterID)

	char := &Character{
		ID:   ch.CharacterID,
		Name: newName.Name,
	}

	log.Println("Saving character:", char)
	mdb := &MongoDBConnection{}
	mdb.session = mdb.GetSession()
	defer mdb.session.Close()
	err := mdb.Save(char)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{"error while saving character:" + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, char)
}

func loadCharacter(c *gin.Context) {

	charID := c.Params.ByName("ID")
	var resultCharacter Character
	log.Println("Looking for character with ID:", charID)

	// config := getConfiguration()
	// storage := getStorage(config.Storage)
	//TODO:Replace this with reflection based on configuration
	mdb := &MongoDBConnection{}
	mdb.session = mdb.GetSession()
	defer mdb.session.Close()
	resultCharacter, err := mdb.Load(charID)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{"Error occured while loading character:" + err.Error()})
		return
	}

	c.JSON(http.StatusOK, resultCharacter)
}
