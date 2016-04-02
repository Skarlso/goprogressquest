package main

import (
	"log"
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
	player, err := mdb.Load(id)
	if err != nil {
		panic(err)
	}

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
		// Check items and equip if better
		// Save player (Progress will not be saved because it's irrelevant.)
		if float64(player.Hp) <= (float64(player.Hp) * 0.25) {
			player.Rest()
		}

		if invetoryIsOverLimit(player) {
			player.SellItems()
		}

		// TODO: Do a switch on possible actions:
		// * Encounter
		// * Discovery
		// * Nothing
		player.Attack(SpawnEnemy(player))

		if player.CurrentXp >= player.NextLevelXp {
			player.LevelUp()
		}

		time.Sleep(time.Millisecond * 500)
	}
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
