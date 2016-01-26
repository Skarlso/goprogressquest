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
	t.SkipNow()
	router := gin.New()
	router.POST("/api/1/start", startAdventure)

	req, _ := http.NewRequest("POST", "/api/1/start", strings.NewReader("{\"name\":\"asdf\"}"))
	req.Header.Add("Content-type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Body.String(), "{\"error\":\"Error occured while loading character:not found\"}\n")
}
