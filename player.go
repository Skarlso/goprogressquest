package main

//Item a representation of an Item and it's properties
type Item struct {
	name   string
	dmg    int
	size   int
	weight int
}

//Inventory holds an endless number of Items
type Inventory struct {
	items []Item
}

//Stats contains a groupped information about stats of a character
type Stats struct {
	str int
	agi int
}

//Character is a player character
type Character struct {
	inventory Inventory
	name      string
	stats     Stats
}
