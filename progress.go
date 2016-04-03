package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

// PointsPerLevel is the maximum points a player can get per level.
const PointsPerLevel = 20

// LevelUp Level up a character.
func (c *Character) LevelUp() {
	// Apply basic character changes first.
	color.Set(color.FgMagenta, color.Bold)
	log.Println("************Player reached a new level!************")
	color.Unset()
	stats := distributePoints()
	c.Stats.Agility += stats[0]
	c.Stats.Intelligence += stats[1]
	c.Stats.Luck += stats[2]
	c.Stats.Perception += stats[3]
	c.Stats.Strenght += stats[4]
	c.Stats.Constitution += stats[5]
	c.MaxHp += 50
	c.Hp = c.MaxHp
	c.CurrentXp = 0

	// Apply calculated changes next.
	c.MaxHp += c.MaxHp / c.Stats.Constitution
	c.NextLevelXp += c.Level * 1000
	c.Level++
	mdb.Update(*c)
	green := color.New(color.FgGreen).SprintFunc()
	log.Printf("Current level is:%s\n", green(c.Level))
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
