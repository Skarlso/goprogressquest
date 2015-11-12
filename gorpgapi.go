package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	http.HandleFunc("/", Index)
	log.Printf("Starting server to listen on port: 8989...")
	http.ListenAndServe(":8989", nil)
}

// The main function which starts the rpg.
func handler() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))
}

//Index Index request handler
func Index(w http.ResponseWriter, r *http.Request) {
	welcome := struct {
		message string
	}{
		"Welcome!",
	}
	json.NewEncoder(w).Encode(welcome)
}
