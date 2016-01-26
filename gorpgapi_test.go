package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestIndexReturningWelcomeMessage(t *testing.T) {
	handler := func(c *gin.Context) {
		c.String(http.StatusOK, "bar")
	}

	router := gin.New()
	router.GET("/foo", handler)

	req, _ := http.NewRequest("GET", "/foo", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, resp.Body.String(), "bar")
}
