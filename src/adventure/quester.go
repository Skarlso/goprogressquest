package adventure

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
