package characters

import (
	"log"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const PointsPerLevel = 20

// Character is a player character.
type Character struct {
	ID          string    `json:"id"`
	Inventory   Inventory `json:"inventory"`
	Body        Body      `json:"body"`
	Name        string    `json:"name"`
	Stats       Stats     `json:"stats"`
	Hp          int       `json:"hp"`
	MaxHp       int       `json:"maxhp"`
	CurrentXp   int       `json:"currentxp"`
	NextLevelXp int       `json:"nextlevelxp"`
	Gold        int       `json:"gold"`
	Level       int       `json:"level"`
	Race        int       `json:"race"`
	Cast        int       `json:"cast"`
}

// Item a representation of an Item and it's properties.
type Item struct {
	Name   string `json:"name"`
	ID     int    `json:"id"`
	Dmg    int    `json:"dmg"`
	Weight int    `json:"weight"`
	Armor  int    `json:"armor"`
	Value  int    `json:"value"`
	Chance int    `json:"chance"`
}

// Inventory holds an endless number of Items
type Inventory struct {
	Items    []Item `json:"items"`
	Capacity int    `json:"capacity"`
}

// Stats contains a groupped information about stats of a character
type Stats struct {
	Strenght     int `json:"strength"`
	Agility      int `json:"agility"`
	Intelligence int `json:"intelligence"`
	Perception   int `json:"perception"`
	Luck         int `json:"luck"`
	Constitution int `json:"consititution"`
}

// Body Represents a body of a Player which defines what he wears,
// Player will always automatically wear the best gear.
type Body struct {
	LRing   Item `json:"lring"`
	RRing   Item `json:"rring"`
	Armor   Item `json:"armor"`
	Head    Item `json:"head"`
	Weapond Item `json:"weapond"`
	Shield  Item `json:"shield"`
}

// Cast the cast of a player, like mage, rouge, warrior...
type Cast struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

// Race the race of the player, like elf, gnome, human, dwarf...
type Race struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

// Rest will Replenish Health.
func (c *Character) Rest() {
	c.Hp = c.MaxHp
	color.Set(color.FgBlue)
	log.Println("Player is fully rested.")
	color.Unset()
	DB.Update(*c)
}

// SellItems will clear the inventory.
func (c *Character) SellItems() {
	for _, v := range c.Inventory.Items {
		c.Gold += v.Value
	}
	yellow := color.New(color.FgYellow).SprintFunc()
	color.Set(color.FgBlue)
	log.Printf("Player has sold all items. Current gold is: %s\n", yellow(c.Gold))
	color.Unset()
	c.Inventory.Items = []Item{}
	DB.Update(*c)
}

// Attack an enemy during an encounter.
func (c *Character) Attack(e Enemy) {
	// Fight until enemy is dead, or player is below 25%.
	red := color.New(color.FgRed, color.Bold).SprintFunc()
	log.Printf("Attacking enemy: %s\n", red(e.Name))
	playerHpLimit := int(float64(c.Hp) * 0.25)
	playerDamage := c.Body.Weapond.Dmg - e.Armor
	if playerDamage <= 0 {
		playerDamage = 1
	}
	enemyDamage := e.Damage - (c.Body.Head.Armor + c.Body.Armor.Armor + c.Body.Shield.Armor + c.Body.LRing.Armor + c.Body.RRing.Armor)
	if enemyDamage <= 0 {
		enemyDamage = 1
	}
	for c.Hp > playerHpLimit && e.Hp > 0 {
		e.Hp -= playerDamage
		c.Hp -= enemyDamage
	}
	if e.Hp <= 0 {
		color.Set(color.FgCyan)
		log.Println("Player won!")
		color.Unset()
		c.CurrentXp += e.Xp
		displayProgressBar(c.CurrentXp, c.NextLevelXp)
		c.awardItems(e)
		DB.Update(*c)
		return
	}
	color.Set(color.FgHiRed)
	log.Println("Enemy won. Player has fled with hp: ", c.Hp)
	color.Unset()
	DB.Update(*c)
}

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
	DB.Update(*c)
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

// awardItems awards the items from a monster based on occurrence chance.
func (c *Character) awardItems(e Enemy) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for _, v := range e.Items {
		ch := r.Intn(100) + 1
		if ch <= v.Chance {
			c.Inventory.Items = append(c.Inventory.Items, v)
		}
	}
}

// displayProgressBar displays a bar representing how much is left to for the character
// to level up.
func displayProgressBar(currXp, nextXp int) {
	progress := "Progress: |"
	top := 100
	dots := top - (currXp / (nextXp / 100))
	hashes := top - dots
	percentage := hashes
	for i := 0; i < top; i++ {
		if hashes > 0 {
			progress += "#"
			hashes--
		} else if dots > 0 {
			progress += "."
			dots--
		}
	}
	progress += "|"
	log.Printf("Progress: %s; Percentage: %d%%", progress, percentage)
}

// checkForBetterItems checks the players inventory for better items to wear.
func (c *Character) checkForBetterItems() {

}
