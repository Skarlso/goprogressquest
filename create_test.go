package main

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

func TestCreateReturnsAnIdAndHash(t *testing.T) {

	mdb = TestDB{}
	router := gin.New()
	router.POST("/"+APIBASE+"/create", create)

	expectedCharacter := Character{}
	expectedCharacter.ID = fmt.Sprintf("%x", sha1.Sum([]byte("asdf")))
	expectedCharacter.Name = "asdf"

	req, _ := http.NewRequest("POST", "/"+APIBASE+"/create", strings.NewReader("{\"name\":\"asdf\"}"))
	req.Header.Set("Content-type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	returnCharacter := Character{}
	json.Unmarshal(resp.Body.Bytes(), &returnCharacter)
	log.Println("Expected:", expectedCharacter)
	log.Println("Actual:", returnCharacter)
	assert.Equal(t, expectedCharacter, returnCharacter)
}

func TestCreateSameCharacterTwice(t *testing.T) {
	t.SkipNow()
}

func TestSavingErrorReturnsProperErrorMessage(t *testing.T) {

	mdb = TestDB{}
	router := gin.New()
	router.POST("/"+APIBASE+"/create", create)

	req, _ := http.NewRequest("POST", "/"+APIBASE+"/create", strings.NewReader("{\"name\":\"save_error\"}"))
	req.Header.Set("Content-type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, "{\"error\":\"error while saving character:error\"}\n", resp.Body.String())
}

func TestMarshallErrorReturnsProperErrorMessage(t *testing.T) {

	mdb = TestDB{}
	router := gin.New()
	router.POST("/"+APIBASE+"/create", create)

	req, _ := http.NewRequest("POST", "/"+APIBASE+"/create", strings.NewReader("invalidJSON"))
	req.Header.Set("Content-type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, "{\"error\":\"error while binding newName:invalid character 'i' looking for beginning of value\"}\n", resp.Body.String())
}

func TestLoadingACharacterWhichWasNotCreated(t *testing.T) {
	mdb = TestDB{}
	router := gin.New()
	router.GET("/"+APIBASE+"/load/:id", loadCharacter)

	req, _ := http.NewRequest("GET", "/"+APIBASE+"/load/not_found", nil)
	req.Header.Set("Content-type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, "{\"error\":\"Error occured while loading character:not found\"}\n", resp.Body.String())
}
