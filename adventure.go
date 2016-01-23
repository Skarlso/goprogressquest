package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

var adventureSignal = make(chan bool, 1)

//TODO: For now, adventuring is saved to a map based on an ID

//StartAdventure starts and adventure in an endless for loop, until a channel signals otherwise
func startAdventure(w http.ResponseWriter, r *http.Request) {
	//First, make it work.
	//second, make it right.
	//Third, make it fast.
	var adventurer struct {
		ID string `json:"id"`
	}

	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&adventurer)
	if err != nil {
		handleError(w, "Error occured while reading Json."+err.Error(), http.StatusBadRequest)
		return
	}

	mdb := MongoDBConnection{}
	mdb.session = mdb.GetSession()
	char, err := mdb.Load(adventurer.ID)
	if err != nil {
		handleError(w, "Error occured while loading character:"+err.Error(), http.StatusBadRequest)
		return
	}

	m := Message{}
	m.Message = "Started adventuring for character: " + char.Name
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)
	//TODO: Next: A way to detect that an adventurer is adventuring.
	go func(name string) {
		stop := false
		for {
			select {
			case stop = <-adventureSignal:
			default:
			}

			if stop {
				log.Println("Stopping adventuring for:", name)
				break
			}

			log.Println("Adventuring...")
			time.Sleep(time.Millisecond * 500)
		}
	}(char.Name)
}

//StopAdventure Stop adventuring
func stopAdventure(w http.ResponseWriter, r *http.Request) {
	//signal channel to stop fight.
	var adventurer struct {
		ID string `json:"id"`
	}

	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&adventurer)
	if err != nil {
		handleError(w, "Error occured while reading Json."+err.Error(), http.StatusBadRequest)
		return
	}

	mdb := MongoDBConnection{}
	mdb.session = mdb.GetSession()
	char, err := mdb.Load(adventurer.ID)
	if err != nil {
		handleError(w, "Error occured while loading character:"+err.Error(), http.StatusBadRequest)
		return
	}

	select {
	case adventureSignal <- true:
	default:
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	m := Message{}
	m.Message = "Stop adventuring signalled for adventurer:" + char.Name
	json.NewEncoder(w).Encode(m)
}
