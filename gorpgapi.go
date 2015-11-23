package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

//APIVERSION Is the current API version
const APIVERSION = "1"

//READLIMIT 1 MiB read limit
const READLIMIT = 1048576

//APIBASE Defines the API base URI
const APIBASE = "/api/" + APIVERSION

// The main function which starts the rpg
func main() {
	handlerChain := alice.New(Logging, PanicHandler)
	router := mux.NewRouter().StrictSlash(true)
	router.Handle(APIBASE+"/", handlerChain.ThenFunc(index)).Methods("GET")
	router.Handle(APIBASE+"/create", handlerChain.ThenFunc(create)).Methods("POST")
	log.Printf("Starting server to listen on port: 8989...")
	log.Fatal(http.ListenAndServe(":8989", router))
}

//index a humble welcome to a new user
func index(w http.ResponseWriter, r *http.Request) {
	m := Message{}
	m.Message = "Welcome to my RPG"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)
}

//create handling the creation of a new character
func create(w http.ResponseWriter, r *http.Request) {
	var newName struct {
		Name string `json:"name"`
	}
	ch := NewCharacter{}
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, READLIMIT))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500) // unprocessable entity
		json.NewEncoder(w).Encode(err)
	}
	if err := r.Body.Close(); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500) // unprocessable entity
		json.NewEncoder(w).Encode(err)
	}
	if err := json.Unmarshal(body, &newName); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(422) // unprocessable entity
		json.NewEncoder(w).Encode(err)
	}

	checkSum := sha1.Sum([]byte(newName.Name))
	ch.CharacterID = fmt.Sprintf("%x", checkSum)
	log.Printf("Created character sha hash: %v", ch.CharacterID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ch)
}
