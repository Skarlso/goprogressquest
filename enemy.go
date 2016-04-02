//enemy a package discribing the properties of an enemy
package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//Enemy represents an enemy combatant.
type Enemy struct {
	Name string
	ID   string
	Race int
	Cast int
	//Items which the player can loot. Will be crossreferenced with Items, from items.json
	Items []Item
	//Gold which the player can loot
	Gold int
	//Xp is calculated based on level and rareness
	Xp int
	//Level is calculated based on the Players level. +-5%
	Level int
	//RarenessLevel is 1-10 where 10 is highly rare
	RarenessLevel int
}

// SpawnEnemy spawns an enemy combatand who's stats are based on the player's character.
// Also, based on RarenessLevel.
func SpawnEnemy(c Character) Enemy {
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

	m.initializeStatsFromJSON()

	return m
}

// MonsterItem is an item in the monsters.json file.
type MonsterItem struct {
	ID int `json:"id"`
}

// Monster is a monster from the monsters.json file.
type Monster struct {
	Name  string        `json:"name"`
	ID    int           `json:"id"`
	Race  int           `json:"race"`
	Cast  int           `json:"cast"`
	Items []MonsterItem `json:"items"`
	Gold  int           `json:"gold"`
	Rare  int           `json:"rare"`
}

// Monsters is a collection of monsters.
type Monsters struct {
	Monster []Monster `json:"monsters"`
}

func (e *Enemy) initializeStatsFromJSON() {

	m := Monsters{}

	file, err := os.Open("monsters.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	data, _ := ioutil.ReadAll(file)

	err = json.Unmarshal(data, &m)
	if err != nil {
		panic(err)
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := r.Intn(len(m.Monster))
	e.Cast = m.Monster[index].Cast
	e.Race = m.Monster[index].Race
	e.RarenessLevel = m.Monster[index].Rare
	e.Gold = m.Monster[index].Gold
	e.ID = strconv.Itoa(m.Monster[index].ID)
	e.Name = m.Monster[index].Name
}
