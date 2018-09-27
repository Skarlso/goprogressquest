package main

import (
	"fmt"
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
		Name string `json:"name"`
	}

	if err := c.BindJSON(&adventurer); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{"error while binding adventurer:" + err.Error()})
		return
	}

	char, err := mdb.Load(adventurer.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{"Error occured while loading character:" + err.Error()})
		return
	}

	adventurersOnQuest.RLock()
	adv := adventurersOnQuest.m[char.Name]
	adventurersOnQuest.RUnlock()

	if adv {
		c.JSON(http.StatusBadRequest, ErrorResponse{"Error occured, adventurer is already adventuring!"})
		return
	}

	go adventuring(char.Name)

	m := Message{}
	m.Message = "Started adventuring for character: " + char.Name
	c.JSON(http.StatusOK, m)
}

func adventuring(name string) {
	adventurersOnQuest.Lock()
	adventurersOnQuest.m[name] = true
	adventurersOnQuest.Unlock()
	stop := false
	player, err := mdb.Load(name)
	if err != nil {
		fmt.Println("could not load player... skipping adventure")
		return
	}

	for {
		select {
		case stop = <-adventureSignal:
		default:
		}

		if stop {
			log.Println("Stopping adventuring for:", name)
			adventurersOnQuest.Lock()
			adventurersOnQuest.m[name] = false
			adventurersOnQuest.Unlock()
			break
		}

		log.Println("Adventuring...")
		if player.Hp <= int((float64(player.MaxHp) * 0.25)) {
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
		player.checkForBetterItems()
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
		Name string `json:"name"`
	}

	if err := c.BindJSON(&adventurer); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{"error while binding adventurer:" + err.Error()})
		return
	}

	char, err := mdb.Load(adventurer.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{"Error occured while loading character:" + err.Error()})
		return
	}

	log.Println("Loaded adventurer:", char)

	adventurersOnQuest.RLock()
	adv := adventurersOnQuest.m[char.Name]
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
