package main

import (
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var adventureSignal = make(chan bool, 1)

// AdventurerOnQuest advneturer On quest with Locking
type AdventurerOnQuest struct {
	m map[string]bool
	sync.RWMutex
}

// TODO: For now, adventuring is saved to a map based on an ID
var adventurersOnQuest = AdventurerOnQuest{m: make(map[string]bool, 0)}

// StartAdventure starts and adventure in an endless for loop, until a channel signals otherwise
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

	adventurersOnQuest.RLock()
	adv := adventurersOnQuest.m[char.ID]
	adventurersOnQuest.RUnlock()

	if adv {
		c.JSON(http.StatusBadRequest, ErrorResponse{"Error occured, adventurer is already adventuring!"})
		return
	}

	go adventuring(char.ID, char.Name)

	m := Message{}
	m.Message = "Started adventuring for character: " + char.Name
	c.JSON(http.StatusOK, m)
}

func adventuring(id string, name string) {
	adventurersOnQuest.Lock()
	adventurersOnQuest.m[id] = true
	adventurersOnQuest.Unlock()
	stop := false
	for {
		select {
		case stop = <-adventureSignal:
		default:
		}

		if stop {
			log.Println("Stopping adventuring for:", name)
			adventurersOnQuest.Lock()
			adventurersOnQuest.m[id] = false
			adventurersOnQuest.Unlock()
			break
		}

		log.Println("Adventuring...")
		// TODO:
		// Get the player
		// Before steps:
		// Low Health? => Rest
		// Inventory full? => Sell
		// Steps:
		// Encounter an enemy
		// Fight ->
		// Low health => Flee && Rest
		// Won -> Avard Xp ->
		// Level up? => Level up
		player, err := mdb.Load(id)
		if err != nil {
			panic(err)
		}
		if float64(player.Hp) < float64(player.Hp)*0.25 {
			player.Rest()
		}

		if invetoryIsOverLimit(player) {
			player.SellItems()
		}

		player.Attack(spawnEnemy(player))

		time.Sleep(time.Millisecond * 500)
	}
}

// spawnEnemy spawns an enemy combatand who's stats are based on the player's character.
func spawnEnemy(c Character) Enemy {
	// Monster Level will be +- 20% of Character Level
	m := Enemy{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	limiter := int(float64(c.Level) * 0.2)
	if limiter <= 0 {
		limiter = 1
	}

	m.Level = (c.Level - limiter) + r.Intn(limiter*2)

	if m.Level < 0 {
		m.Level = 0
	}
	return m
}

func invetoryIsOverLimit(c Character) bool {
	currCap := 0
	for _, v := range c.Inventory.Items {
		currCap += v.Weight
	}

	if currCap >= c.Inventory.Capacity {
		return true
	}

	return false
}

// StopAdventure Stop adventuring
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

	adventurersOnQuest.RLock()
	adv := adventurersOnQuest.m[char.ID]
	adventurersOnQuest.RUnlock()
	if !adv {
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
