//enemy a package discribing the properties of an enemy
package main

//Enemy discribes an enemy combatant.
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
