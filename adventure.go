package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

var adventureSignal = make(chan bool)

//StartAdventure starts and adventure in an endless for loop, until a channel signals otherwise
func startAdventure(w http.ResponseWriter, r *http.Request) {
	//First, make it work.
	//second, make it right.
	//Third, make it fast.
	var adventurer struct {
		Name string `json:"name"`
	}

	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&adventurer)

	if err != nil {
		handleError(w, "Error occured while reading Json."+err.Error(), http.StatusBadRequest)
	}

	m := Message{}
	m.Message = "Started adventuring for character: " + adventurer.Name
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)
	go func() {
		// w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusOK)
		for {
			select {
			case <-adventureSignal:
				log.Println("Stopping adventuring...")
				return
			default:
				log.Println("Adventuring...")
				time.Sleep(time.Millisecond * 500)
			}
		}
	}()
}

//StopAdventure Stop adventuring
func stopAdventure(w http.ResponseWriter, r *http.Request) {
	//signal channel to stop fight.
	var adventurer struct {
		Name string `json:"name"`
	}

	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&adventurer)

	if err != nil {
		handleError(w, "Error occured while reading Json."+err.Error(), http.StatusBadRequest)
	}
	adventureSignal <- true
	// select {
	// case adventureSignal <- true:
	// default:
	// }
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	m := Message{}
	m.Message = "Stop adventuring signalled for adventurer:" + adventurer.Name
	json.NewEncoder(w).Encode(m)
}
