package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

//APIVERSION Is the current API version
const APIVERSION = "1"

//APIBASE Defines the API base URI
const APIBASE = "api/" + APIVERSION

//Config global configuration of the application
type Config struct {
	Storage string `json:"Storage"`
}

func getConfiguration() (c Config) {
	dat, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(dat, &c); err != nil {
		panic(err)
	}

	return
}

// The main function which starts the rpg
func main() {
	router := gin.Default()
	v1 := router.Group(APIBASE)
	{
		v1.GET("/", index)
		v1.POST("/create", create)
		v1.GET("/load/:id", loadCharacter)
		v1.POST("/start", startAdventure)
		v1.POST("/stop", stopAdventure)
	}
	router.Run(":8989")
}

//index a humble welcome to a new user
func index(c *gin.Context) {
	m := Message{}
	m.Message = "Welcome to my RPG"
	c.JSON(http.StatusOK, m)
}
