package main

import "fmt"

// Starting mongodb -> mongod --config /usr/local/etc/mongod.conf --fork

// TestDB Encapsulates a connection to a database
type TestDB struct {
}

// Save will save a player using mongodb as a storage medium
func (tdb TestDB) Save(ch Character) error {
	if ch.Name == "save_error" {
		return fmt.Errorf("error")
	}
	return nil
}

// Load will load the player using mongodb as a storage medium
func (tdb TestDB) Load(Name string) (result Character, err error) {
	if Name == "not_found" {
		return Character{}, fmt.Errorf("not found")
	}
	return Character{ID: "asdf", Name: Name}, nil
}

// Update update a character
func (tdb TestDB) Update(c Character) error {
	return nil
}
