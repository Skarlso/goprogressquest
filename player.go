package main

//Item a representation of an Item and it's properties
type Item struct {
	Name   string
	Dmg    int
	Size   int
	Weight int
}

//Inventory holds an endless number of Items
type Inventory struct {
	Items []Item
}

//Stats contains a groupped information about stats of a character
type Stats struct {
	Strenght     int
	Agility      int
	Intelligence int
	Perception   int
	Luck         int
}

//Body Represents a body of a Player which defines what he wears
type Body struct {
	LRing   Item
	RRing   Item
	Armor   Item
	Head    Item
	Weapond Item
	Shield  Item
}

//Cast the cast of a player, like mage, rouge, warrior...
type Cast struct {
	Name string
	ID   int
}

//Race the race of the player, like elf, gnome, human, dwarf...
type Race struct {
	Name string
	ID   int
}

//Character is a player character
type Character struct {
	ID        string
	Inventory Inventory
	Body      Body
	Name      string
	Stats     Stats
	Gold      int
	Level     int
	Race      int
	Cast      int
}
