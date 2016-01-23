package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//create handling the creation of a new character
//curl -H "Content-Type: application/json" -X POST -d '{"name":"asdf"}' http://localhost:8989
func create(w http.ResponseWriter, r *http.Request) {
	var newName struct {
		Name string `json:"name"`
	}
	ch := NewCharacter{}
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newName)

	if err != nil {
		handleError(w, "Error occured while reading Json."+err.Error(), http.StatusBadRequest)
		return
	}

	checkSum := sha1.Sum([]byte(newName.Name))
	ch.CharacterID = fmt.Sprintf("%x", checkSum)
	log.Printf("Created character sha hash: %v", ch.CharacterID)

	char := &Character{
		ID:   ch.CharacterID,
		Name: newName.Name,
	}

	log.Println("Saving character:", char)
	mdb := &MongoDBConnection{}
	mdb.session = mdb.GetSession()
	defer mdb.session.Close()
	err = mdb.Save(char)
	if err != nil {
		handleError(w, "Error occured while saving character."+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ch)
}

func loadCharacter(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	charID := vars["ID"]
	var resultCharacter Character
	log.Println("Looking for character with ID:", charID)

	// config := getConfiguration()
	// storage := getStorage(config.Storage)
	//TODO:Replace this with reflection based on configuration
	mdb := &MongoDBConnection{}
	mdb.session = mdb.GetSession()
	defer mdb.session.Close()
	resultCharacter, err := mdb.Load(charID)
	if err != nil {
		handleError(w, "Error occured while loading character:"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	//Not handling error cases yet when the Character could not be retrieved
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(resultCharacter)
}
