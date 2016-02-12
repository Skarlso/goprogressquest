package main

import (
	"math/rand"
	"time"
)

// PointsPerLevel is the maximum points a player can get per level.
const PointsPerLevel = 20

// LevelUp Level up a character.
func (c *Character) LevelUp() {
	// Apply basic character changes first.
	stats := distributePoints()
	c.Stats.Agility += stats[0]
	c.Stats.Intelligence += stats[1]
	c.Stats.Luck += stats[2]
	c.Stats.Perception += stats[3]
	c.Stats.Strenght += stats[4]
	c.Stats.Constitution += stats[5]
	c.Hp += 50
	c.CurrentXp = 0

	// Apply calculated changes next.
	c.Hp += c.Hp / c.Stats.Constitution
	c.NextLevelXp += c.Level * 1000
	c.Level++
}

func distributePoints() []int {
	currentPoints := PointsPerLevel
	var stats []int
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 6; i >= 1; i-- {
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