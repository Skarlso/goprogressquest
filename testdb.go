package main

//Starting mongodb -> mongod --config /usr/local/etc/mongod.conf --fork

//TestDB Encapsulates a connection to a database
type TestDB struct {
}

//Save will save a player using mongodb as a storage medium
func (tdb TestDB) Save(ch Character) error {
	return nil
}

//Load will load the player using mongodb as a storage medium
func (tdb TestDB) Load(ID string) (result Character, err error) {
	return Character{}, nil
}
