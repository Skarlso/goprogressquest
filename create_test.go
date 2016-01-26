package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateReturnsAnIdAndHash(t *testing.T) {
	router := gin.New()
	router.POST("/api/1/create", create)

	expectedCharacter := Character{}
	expectedCharacter.ID = fmt.Sprintf("%x", sha1.Sum([]byte("asdf")))
	expectedCharacter.Name = "asdf"

	req, _ := http.NewRequest("POST", "/api/1/create", strings.NewReader("{\"name\":\"asdf\"}"))
	req.Header.Set("Content-type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	returnCharacter := Character{}

	json.Unmarshal(resp.Body.Bytes(), &returnCharacter)

	assert.Equal(t, expectedCharacter, returnCharacter)
}

func TestCreateSameCharacterTwice(t *testing.T) {
	t.SkipNow()
}
