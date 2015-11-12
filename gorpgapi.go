package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// The main function which starts the rpg.
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	log.Printf("Starting server to listen on port: 8989...")
	log.Fatal(http.ListenAndServe(":8989", router))
}

//Index Index request handler
func Index(w http.ResponseWriter, r *http.Request) {
	welcome := struct {
		//Message in order for it to be unmarshalled it must be exported
		Message string `json:"message"`
	}{
		"Welcome!",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(welcome)
}
