package main

import (
	"encoding/json"
	"net/http"
	"time"
)

var adventureSignal = make(chan bool)

//StartAdventure starts and adventure in an endless for loop, until a channel signals otherwise
func StartAdventure(w http.ResponseWriter, r *http.Request) {
	//First, make it work.
	//second, make it right.
	//Third, make it fast.
	m := Message{}
	m.Message = "Started adventuring"
	go func() {
		for {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(m)
			if <-adventureSignal {
				m = Message{}
				m.Message = "Stopped adventuring"
				json.NewEncoder(w).Encode(m)
				break
			}
			time.Sleep(time.Millisecond * 500)
		}
	}()
}

//StopAdventure Stop adventuring
func StopAdventure(w http.ResponseWriter, r *http.Request) {
	//signal channel to stop fight.
	adventureSignal <- true
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	m := Message{}
	m.Message = "Stop adventuring signalled."
	json.NewEncoder(w).Encode(m)
}
