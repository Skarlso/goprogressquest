package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAdventureReturningErrorOnPlayerWhichIsNotCreatedOnStart(t *testing.T) {
	mdb = TestDB{}
	router := gin.New()
	router.POST("/"+APIBASE+"/start", startAdventure)

	req, _ := http.NewRequest("POST", "/"+APIBASE+"/start", strings.NewReader("{\"name\":\"not_found\"}"))
	req.Header.Add("Content-type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, "{\"error\":\"Error occured while loading character:not found\"}\n", resp.Body.String())
}

func TestAdventureReturningErrorOnPlayerWhichIsNotCreatedOnStop(t *testing.T) {
	mdb = TestDB{}
	router := gin.New()
	router.POST("/"+APIBASE+"/stop", stopAdventure)

	req, _ := http.NewRequest("POST", "/"+APIBASE+"/stop", strings.NewReader("{\"name\":\"not_found\"}"))
	req.Header.Add("Content-type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, "{\"error\":\"Error occured while loading character:not found\"}\n", resp.Body.String())
}

func TestStartingAdventuringForPlayerWhoIsAdventuring(t *testing.T) {
	mdb = TestDB{}
	adventurersOnQuest.Lock()
	adventurersOnQuest.m["onquest"] = true
	adventurersOnQuest.Unlock()
	router := gin.New()
	router.POST("/"+APIBASE+"/start", startAdventure)

	req, _ := http.NewRequest("POST", "/"+APIBASE+"/start", strings.NewReader("{\"name\":\"onquest\"}"))
	req.Header.Add("Content-type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, "{\"error\":\"Error occured, adventurer is already adventuring!\"}\n", resp.Body.String())
}

func TestErrorWhileBindingAdventurerOnStart(t *testing.T) {
	mdb = TestDB{}
	router := gin.New()
	router.POST("/"+APIBASE+"/start", startAdventure)

	req, _ := http.NewRequest("POST", "/"+APIBASE+"/start", strings.NewReader("invalid"))
	req.Header.Add("Content-type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, "{\"error\":\"error while binding adventurer:invalid character 'i' looking for beginning of value\"}\n", resp.Body.String())
}

func TestStopAdventuringForACharacterWhichIsNotAdventuring(t *testing.T) {
	mdb = TestDB{}
	router := gin.New()
	router.POST("/"+APIBASE+"/stop", stopAdventure)

	req, _ := http.NewRequest("POST", "/"+APIBASE+"/stop", strings.NewReader("{\"name\":\"notonquest\"}"))
	req.Header.Add("Content-type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, "{\"error\":\"Error occured, adventurer is not adventuring!\"}\n", resp.Body.String())
}

func TestStopAdventuringInvalidJSON(t *testing.T) {
	mdb = TestDB{}
	router := gin.New()
	router.POST("/"+APIBASE+"/stop", stopAdventure)

	req, _ := http.NewRequest("POST", "/"+APIBASE+"/stop", strings.NewReader("invalid"))
	req.Header.Add("Content-type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, "{\"error\":\"error while binding adventurer:invalid character 'i' looking for beginning of value\"}\n", resp.Body.String())
}

func TestStartAdventuringForExistingPlayer(t *testing.T) {
	mdb = TestDB{}
	router := gin.New()
	router.POST("/"+APIBASE+"/start", startAdventure)

	req, _ := http.NewRequest("POST", "/"+APIBASE+"/start", strings.NewReader("{\"name\":\"quester\"}"))
	req.Header.Add("Content-type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, "{\"message\":\"Started adventuring for character: quester\"}\n", resp.Body.String())
}

func TestStopAdventuringForAdventurerWhoIsAdventuring(t *testing.T) {
	mdb = TestDB{}
	router := gin.New()
	adventurersOnQuest.Lock()
	adventurersOnQuest.m["quester"] = true
	adventurersOnQuest.Unlock()
	router.POST("/"+APIBASE+"/stop", stopAdventure)
	req, _ := http.NewRequest("POST", "/"+APIBASE+"/stop", strings.NewReader("{\"name\":\"quester\"}"))
	req.Header.Add("Content-type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, "{\"message\":\"Stop adventuring for character: quester\"}\n", resp.Body.String())
	//Also assert that our adventureSignal was fired of
	assert.Equal(t, 1, len(adventureSignal))
}

func TestAdventuring(t *testing.T) {
	t.SkipNow()
	go adventuring("name")
	// fmt.Println(adventurersOnQuest)
	// adv := adventurersOnQuest["id"]
	adventureSignal <- true
	// assert.Equal(t, adv, true)
}
