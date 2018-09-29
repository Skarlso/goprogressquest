package adventure

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Skarlso/goprogressquest/src/characters"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestDB Encapsulates a connection to a database
type TestDB struct {
}

// Save will save a player using mongodb as a storage medium
func (tdb TestDB) Save(ch characters.Character) error {
	if ch.Name == "save_error" {
		return fmt.Errorf("error")
	}
	return nil
}

// Load will load the player using mongodb as a storage medium
func (tdb TestDB) Load(Name string) (result characters.Character, err error) {
	if Name == "not_found" {
		return characters.Character{}, fmt.Errorf("not found")
	}
	return characters.Character{ID: "asdf", Name: Name}, nil
}

// Update update a character
func (tdb TestDB) Update(c characters.Character) error {
	return nil
}

func TestAdventureReturningErrorOnPlayerWhichIsNotCreatedOnStart(t *testing.T) {
	characters.DB = TestDB{}
	router := gin.New()
	router.POST("/api/1/start", StartAdventure)

	req, _ := http.NewRequest("POST", "/api/1/start", strings.NewReader("{\"name\":\"not_found\"}"))
	req.Header.Add("Content-type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, "{\"error\":\"Error occured while loading character:not found\"}", resp.Body.String())
}

func TestAdventureReturningErrorOnPlayerWhichIsNotCreatedOnStop(t *testing.T) {
	characters.DB = TestDB{}
	router := gin.New()
	router.POST("/api/1/stop", StopAdventure)

	req, _ := http.NewRequest("POST", "/api/1/stop", strings.NewReader("{\"name\":\"not_found\"}"))
	req.Header.Add("Content-type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, "{\"error\":\"Error occured while loading character:not found\"}", resp.Body.String())
}

func TestStartingAdventuringForPlayerWhoIsAdventuring(t *testing.T) {
	characters.DB = TestDB{}
	adventurersOnQuest.Lock()
	adventurersOnQuest.m["onquest"] = true
	adventurersOnQuest.Unlock()
	router := gin.New()
	router.POST("/api/1/start", StartAdventure)

	req, _ := http.NewRequest("POST", "/api/1/start", strings.NewReader("{\"name\":\"onquest\"}"))
	req.Header.Add("Content-type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, "{\"error\":\"Error occured, adventurer is already adventuring!\"}", resp.Body.String())
}

func TestErrorWhileBindingAdventurerOnStart(t *testing.T) {
	characters.DB = TestDB{}
	router := gin.New()
	router.POST("/api/1/start", StartAdventure)

	req, _ := http.NewRequest("POST", "/api/1/start", strings.NewReader("invalid"))
	req.Header.Add("Content-type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, "{\"error\":\"error while binding adventurer:invalid character 'i' looking for beginning of value\"}", resp.Body.String())
}

func TestStopAdventuringForACharacterWhichIsNotAdventuring(t *testing.T) {
	characters.DB = TestDB{}
	router := gin.New()
	router.POST("/api/1/stop", StopAdventure)

	req, _ := http.NewRequest("POST", "/api/1/stop", strings.NewReader("{\"name\":\"notonquest\"}"))
	req.Header.Add("Content-type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, "{\"error\":\"Error occured, adventurer is not adventuring!\"}", resp.Body.String())
}

func TestStopAdventuringInvalidJSON(t *testing.T) {
	characters.DB = TestDB{}
	router := gin.New()
	router.POST("/api/1/stop", StopAdventure)

	req, _ := http.NewRequest("POST", "/api/1/stop", strings.NewReader("invalid"))
	req.Header.Add("Content-type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, "{\"error\":\"error while binding adventurer:invalid character 'i' looking for beginning of value\"}", resp.Body.String())
}

func TestStartAdventuringForExistingPlayer(t *testing.T) {
	characters.DB = TestDB{}
	router := gin.New()
	router.POST("/api/1/start", StartAdventure)

	req, _ := http.NewRequest("POST", "/api/1/start", strings.NewReader("{\"name\":\"quester\"}"))
	req.Header.Add("Content-type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, "{\"message\":\"Started adventuring for character: quester\"}", resp.Body.String())
}

func TestStopAdventuringForAdventurerWhoIsAdventuring(t *testing.T) {
	characters.DB = TestDB{}
	router := gin.New()
	adventurersOnQuest.Lock()
	adventurersOnQuest.m["quester"] = true
	adventurersOnQuest.Unlock()
	router.POST("/api/1/stop", StopAdventure)
	req, _ := http.NewRequest("POST", "/api/1/stop", strings.NewReader("{\"name\":\"quester\"}"))
	req.Header.Add("Content-type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, "{\"message\":\"Stop adventuring for character: quester\"}", resp.Body.String())
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
