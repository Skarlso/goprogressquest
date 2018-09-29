package characters

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestDB Encapsulates a connection to a database
type TestDB struct {
}

// Save will save a player using mongodb as a storage medium
func (tdb TestDB) Save(ch Character) error {
	if ch.Name == "save_error" {
		return fmt.Errorf("error")
	}
	return nil
}

// Load will load the player using mongodb as a storage medium
func (tdb TestDB) Load(Name string) (result Character, err error) {
	if Name == "not_found" {
		return Character{}, fmt.Errorf("not found")
	}
	return Character{ID: "asdf", Name: Name}, nil
}

// Update update a character
func (tdb TestDB) Update(c Character) error {
	return nil
}

func TestCreateReturnsAnIdAndHash(t *testing.T) {
	DB = TestDB{}
	router := gin.New()
	router.POST("/api/1/create", Create)

	expectedCharacter := Character{}
	expectedCharacter.ID = fmt.Sprintf("%x", sha1.Sum([]byte("asdf")))
	expectedCharacter.Name = "asdf"

	req, _ := http.NewRequest("POST", "/api/1/create", strings.NewReader("{\"name\":\"asdf\"}"))
	req.Header.Set("Content-type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	returnCharacter := Character{}
	json.Unmarshal(resp.Body.Bytes(), &returnCharacter)
	log.Println("Expected:", expectedCharacter)
	log.Println("Actual:", returnCharacter)
	assert.Equal(t, expectedCharacter.ID, returnCharacter.ID)
	assert.Equal(t, expectedCharacter.Name, returnCharacter.Name)
}

func TestCreateSameCharacterTwice(t *testing.T) {
	t.SkipNow()
}

func TestSavingErrorReturnsProperErrorMessage(t *testing.T) {
	DB = TestDB{}
	router := gin.New()
	router.POST("/api/1/create", Create)

	req, _ := http.NewRequest("POST", "/api/1/create", strings.NewReader("{\"name\":\"save_error\"}"))
	req.Header.Set("Content-type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, "{\"error\":\"error while saving character:error\"}", resp.Body.String())
}

func TestMarshallErrorReturnsProperErrorMessage(t *testing.T) {
	DB = TestDB{}
	router := gin.New()
	router.POST("/api/1/create", Create)

	req, _ := http.NewRequest("POST", "/api/1/create", strings.NewReader("invalidJSON"))
	req.Header.Set("Content-type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, "{\"error\":\"error while binding newName:invalid character 'i' looking for beginning of value\"}", resp.Body.String())
}

func TestLoadingACharacterWhichWasNotCreated(t *testing.T) {
	DB = TestDB{}
	router := gin.New()
	router.GET("/api/1/load/:name", LoadCharacter)

	req, _ := http.NewRequest("GET", "/api/1/load/not_found", nil)
	req.Header.Set("Content-type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, "{\"error\":\"Error occured while loading character:not found\"}", resp.Body.String())
}

func TestLoadingCharacter(t *testing.T) {
	DB = TestDB{}
	router := gin.New()
	router.GET("/api/1/load/:name", LoadCharacter)

	expectedCharacter := Character{}
	expectedCharacter.ID = "asdf"
	expectedCharacter.Name = "asdf"

	req, _ := http.NewRequest("GET", "/api/1/load/asdf", nil)
	req.Header.Set("Content-type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	returnCharacter := Character{}
	json.Unmarshal(resp.Body.Bytes(), &returnCharacter)
	// log.Println("Expected:", expectedCharacter)
	// log.Println("Actual:", returnCharacter)
	assert.Equal(t, expectedCharacter, returnCharacter)
}

func TestCharacterCreation(t *testing.T) {
	t.SkipNow()
}

func TestGettingCasts(t *testing.T) {
	t.SkipNow()
}

func TestGettingRaces(t *testing.T) {
	t.SkipNow()
}
