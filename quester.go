package main

const (
	//DISCOVERY Find something. Item, money.
	DISCOVERY = 1 << iota
	//ENCOUNTER Meet an enemy
	ENCOUNTER
	//NEUTRAL Nothing
	NEUTRAL
)

const (
	ELF = 1 << iota
	HUMAN
	DWARF
	FAIRY
	HOBBIT
)

const (
	WARRIOR = 1 << iota
	MAGE
	ROUGE
	BLACKSMITH
	WINDRUNNER
	CABLEGUY
)

//EventType Type of an Event
type EventType struct {
}

//Event a event
type Event struct {
	eType EventType
}
