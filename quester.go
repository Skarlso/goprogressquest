package main

import (
	"math/rand"
	"time"
)

const (
	//DISCOVERY Find something. Item, money.
	DISCOVERY = 1 << iota
	//ENCOUNTER Meet an enemy
	ENCOUNTER
	//NEUTRAL Nothing
	NEUTRAL
)

// EventType Type of an Event
type EventType struct {
}

// Event a event
type Event struct {
	eType EventType
}

// EncounterEnemy defines an event when the player encounters an enemy and has to fight.
func EncounterEnemy() {

}

// spawnEnemy spawns an enemy combatand who's stats are based on the player's character.
func (c Character) spawnEnemy() Enemy {
	// Monster Level will be +- 20% of Character Level
	m := Enemy{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	limiter := int(float64(c.Level) * 0.2)
	m.Level = (c.Level - limiter) + r.Intn(limiter*2)

	if m.Level < 0 {
		m.Level = 0
	}
	return m
}
