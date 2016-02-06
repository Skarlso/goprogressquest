package main

import (
	"math/rand"
	"time"
)

// PointsPerLevel is the maximum points a player can get per level.
const PointsPerLevel = 15

// LevelUp Level up a character.
func (c *Character) LevelUp() {

}

func distributePoints() []int {
	currentPoints := PointsPerLevel
	var stats []int
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 5; i >= 1; i-- {
		stat := r.Intn(currentPoints-i) + 1
		stats = append(stats, stat)
		currentPoints -= stat
	}
	if currentPoints > 0 {
		stats[len(stats)-1] += currentPoints
	}

	return stats
}

// Attack an enemy during an encounter.
func (c *Character) Attack(e Enemy) {

}

// Shop spend money.
func (c *Character) Shop() {

}

// Progress general progress.
func (c *Character) Progress() {

}
