package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestIndexReturningWelcomeMessage(t *testing.T) {
	indexHandler := index
	router := gin.New()
	router.GET("/", indexHandler)

	req, _ := http.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, "{\"message\":\"Welcome to my RPG\"}\n", resp.Body.String())
}

func TestInvalidAPICallReturnsNotFound(t *testing.T) {
	t.SkipNow()
}

func TestGetConfiguration(t *testing.T) {
	c := getConfiguration()

	assert.NotEmpty(t, c)
}
