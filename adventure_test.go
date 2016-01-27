package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAdventureReturningErrorOnPlayerWhichIsNotCreated(t *testing.T) {
	mdb = TestDB{}
	router := gin.New()
	router.POST("/"+APIBASE+"/start", startAdventure)

	req, _ := http.NewRequest("POST", "/"+APIBASE+"/start", strings.NewReader("{\"id\":\"not_found\"}"))
	req.Header.Add("Content-type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Body.String(), "{\"error\":\"Error occured while loading character:not found\"}\n")
}

func TestStartingAdventuringForPlayerWhoIsAdventuring(t *testing.T) {
	mdb = TestDB{}
	adventurersOnQuest["onquest"] = true
	router := gin.New()
	router.POST("/"+APIBASE+"/start", startAdventure)

	req, _ := http.NewRequest("POST", "/"+APIBASE+"/start", strings.NewReader("{\"id\":\"onquest\"}"))
	req.Header.Add("Content-type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Body.String(), "{\"error\":\"Error occured, adventurer is already adventuring!\"}\n")
}

func TestErrorWhileBindingAdventurer(t *testing.T) {
	mdb = TestDB{}
	router := gin.New()
	router.POST("/"+APIBASE+"/start", startAdventure)

	req, _ := http.NewRequest("POST", "/"+APIBASE+"/start", strings.NewReader("invalid"))
	req.Header.Add("Content-type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Body.String(), "{\"error\":\"error while binding adventurer:invalid character 'i' looking for beginning of value\"}\n")
}

func TestStopAdventuringForACharacterWhichIsNotAdventuring(t *testing.T) {
	mdb = TestDB{}
	router := gin.New()
	router.POST("/"+APIBASE+"/stop", stopAdventure)

	req, _ := http.NewRequest("POST", "/"+APIBASE+"/stop", strings.NewReader("{\"id\":\"notonquest\"}"))
	req.Header.Add("Content-type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Body.String(), "{\"error\":\"Error occured, adventurer is not adventuring!\"}\n")
}

func TestStopAdventuringInvalidJSON(t *testing.T) {
	mdb = TestDB{}
	router := gin.New()
	router.POST("/"+APIBASE+"/stop", stopAdventure)

	req, _ := http.NewRequest("POST", "/"+APIBASE+"/stop", strings.NewReader("invalid"))
	req.Header.Add("Content-type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Body.String(), "{\"error\":\"error while binding adventurer:invalid character 'i' looking for beginning of value\"}\n")
}
