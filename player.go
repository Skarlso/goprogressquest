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
	Str int
	Agi int
	In  int
	Per int
	Chr int
	Lck int
}

//Character is a player character
type Character struct {
	Inventory Inventory
	Name      string ``
	Stats     Stats
	ID        string
	Gold      int
}
