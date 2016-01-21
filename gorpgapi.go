package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/mgo.v2"

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
	router.Handle(APIBASE+"/start", handlerChain.ThenFunc(StartAdventure)).Methods("POST")
	router.Handle(APIBASE+"/stop", handlerChain.ThenFunc(StopAdventure)).Methods("POST")
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
//curl -H "Content-Type: application/json" -X POST -d '{"name":"asdf"}' http://localhost:8989
func create(w http.ResponseWriter, r *http.Request) {
	var newName struct {
		Name string `json:"name"`
	}
	ch := NewCharacter{}
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, READLIMIT))
	if err != nil {
		handleError(w, "Error occured while reading the body:")
		return
	}
	if err := r.Body.Close(); err != nil {
		handleError(w, fmt.Sprintf("Error occured while closing the body: %v", err))
		return
	}
	if err := json.Unmarshal(body, &newName); err != nil {
		handleError(w, fmt.Sprintf("Error occured while formatting request: %v", err))
		return
	}

	checkSum := sha1.Sum([]byte(newName.Name))
	ch.CharacterID = fmt.Sprintf("%x", checkSum)
	log.Printf("Created character sha hash: %v", ch.CharacterID)

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//TODO:Replace this with reflection based on configuration
	mdb := MongoDBConnection{}
	mdb.session = session

	// var mongoChar = struct {
	//
	// }

	char := Character{id: ch.CharacterID, name: newName.Name}

	log.Println("Saving character:", char)
	mdb.Save(char)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ch)
}

func loadCharacter(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	charID := vars["ID"]
	var resultCharacter Character
	log.Println("Looking for character with ID:", charID)
	//TODO: I need to abstract this out more into a global variable which is a Storage
	//that has a connection setup in the init()
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	// config := getConfiguration()
	// storage := getStorage(config.Storage)
	//TODO:Replace this with reflection based on configuration
	mdb := MongoDBConnection{}
	mdb.session = session
	resultCharacter = mdb.Load(charID)

	w.Header().Set("Content-Type", "application/json")
	//Not handling error cases yet when the Character could not be retrieved
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(resultCharacter)
}

//getStorage will be replaced by some kind of reflection mechanism
// func getStorage(storage string) (st Storage) {
// 	switch storage {
// 	case "MongoDBConnection":
// 		st = MongoDBConnection{}
// 	}
// 	return
// }

func handleError(w http.ResponseWriter, s string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500) // unprocessable entity
	errorResponse := ErrorResponse{}
	errorResponse.ErrorMessage = fmt.Sprintf(s)
	json.NewEncoder(w).Encode(errorResponse)
}
