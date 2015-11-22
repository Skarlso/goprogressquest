package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

//APIVERSION Is the current API version
const APIVERSION = "1"

// The main function which starts the rpg
func main() {
	handlerChain := alice.New(Logging, PanicHandler)
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/api/"+APIVERSION+"/", handlerChain.ThenFunc(index))
	router.Handle("/api/"+APIVERSION+"/create", handlerChain.ThenFunc(create))
	log.Printf("Starting server to listen on port: 8989...")
	log.Fatal(http.ListenAndServe(":8989", router))
}

//index Index request handler
func index(w http.ResponseWriter, r *http.Request) {
	m := Message{}
	m.Message = "Welcome to my RPG"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)
}

//register handling the creation of a new character
func create(w http.ResponseWriter, r *http.Request) {
	ch := NewCharacter{}
	ch.CharacterID = "Some Random Generated ID"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ch)
}
