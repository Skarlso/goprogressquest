package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var adventureSignal = make(chan bool, 1)

//TODO: For now, adventuring is saved to a map based on an ID
var adventurersOnQuest = make(map[string]bool, 0)

//StartAdventure starts and adventure in an endless for loop, until a channel signals otherwise
func startAdventure(c *gin.Context) {

	var adventurer struct {
		ID string `json:"id"`
	}

	if err := c.BindJSON(&adventurer); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{"error while binding adventurer:" + err.Error()})
		return
	}

	char, err := mdb.Load(adventurer.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{"Error occured while loading character:" + err.Error()})
		return
	}

	if adventurersOnQuest[char.ID] {
		c.JSON(http.StatusBadRequest, ErrorResponse{"Error occured, adventurer is already adventuring!"})
		return
	}

	go adventuring(char.ID, char.Name)

	m := Message{}
	m.Message = "Started adventuring for character: " + char.Name
	c.JSON(http.StatusOK, m)
}

func adventuring(id string, name string) {
	adventurersOnQuest[id] = true
	stop := false
	for {
		select {
		case stop = <-adventureSignal:
		default:
		}

		if stop {
			log.Println("Stopping adventuring for:", name)
			adventurersOnQuest[id] = false
			break
		}

		log.Println("Adventuring...")
		time.Sleep(time.Millisecond * 500)
	}
}

//StopAdventure Stop adventuring
func stopAdventure(c *gin.Context) {
	//signal channel to stop fight.
	var adventurer struct {
		ID string `json:"id"`
	}

	if err := c.BindJSON(&adventurer); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{"error while binding adventurer:" + err.Error()})
		return
	}

	char, err := mdb.Load(adventurer.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{"Error occured while loading character:" + err.Error()})
		return
	}

	log.Println("Loaded adventurer:", char)
	log.Println("Adventurer is on questing?", adventurersOnQuest[char.ID])

	if !adventurersOnQuest[char.ID] {
		c.JSON(http.StatusBadRequest, ErrorResponse{"Error occured, adventurer is not adventuring!"})
		return
	}
	select {
	case adventureSignal <- true:
	default:
	}

	m := Message{}
	m.Message = "Stop adventuring for character: " + char.Name
	c.JSON(http.StatusOK, m)
}
