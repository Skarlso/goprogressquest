package main

import (
	"net/http"

	"github.com/Skarlso/goprogressquest/src/adventure"
	"github.com/Skarlso/goprogressquest/src/characters"
	"github.com/Skarlso/goprogressquest/src/responsetypes"
	"github.com/gin-gonic/gin"
)

//APIVERSION Is the current API version
const APIVERSION = "1"

//APIBASE Defines the API base URI
const APIBASE = "api/" + APIVERSION

// The main function which starts the rpg
func main() {
	router := gin.Default()
	v1 := router.Group(APIBASE)
	{
		v1.GET("/", index)
		v1.POST("/create", characters.Create)
		v1.GET("/load/:name", characters.LoadCharacter)
		v1.POST("/start", adventure.StartAdventure)
		v1.POST("/stop", adventure.StopAdventure)
	}
	router.Run(":8989")
}

//index a humble welcome to a new user
func index(c *gin.Context) {
	m := responsetypes.Message{}
	m.Message = "Welcome to my RPG"
	c.JSON(http.StatusOK, m)
}
