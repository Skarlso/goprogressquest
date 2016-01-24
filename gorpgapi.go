package main

import (
	"encoding/json"
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

//Config global configuration of the application
type Config struct {
	Storage string `json:"Storage"`
}

func getConfiguration() (c Config) {
	dat, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(dat, &c); err != nil {
		panic(err)
	}

	return
}

// The main function which starts the rpg
func main() {
	handlerChain := alice.New(Logging, PanicHandler)
	router := mux.NewRouter().StrictSlash(true)
	router.Handle(APIBASE+"/", handlerChain.ThenFunc(index)).Methods("GET")
	router.Handle(APIBASE+"/create", handlerChain.ThenFunc(create)).Methods("POST")
	router.Handle(APIBASE+"/load/{ID}", handlerChain.ThenFunc(loadCharacter)).Methods("GET")
	router.Handle(APIBASE+"/start", handlerChain.ThenFunc(startAdventure)).Methods("POST")
	router.Handle(APIBASE+"/stop", handlerChain.ThenFunc(stopAdventure)).Methods("POST")
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

func handleError(w http.ResponseWriter, s string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	errorResponse := ErrorResponse{}
	errorResponse.ErrorMessage = s
	json.NewEncoder(w).Encode(errorResponse)
}
