package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

//APIVERSION Is the current API version
const APIVERSION = "1"

//APIBASE Defines the API base URI
const APIBASE = "api/" + APIVERSION

var mdb Storage
var config Config

// ItemsMap contains all the items from items.json file
var ItemsMap map[int]Item

//Config global configuration of the application
type Config struct {
	Storage string `json:"Storage"`
	DBURL   string `json:"DBURL"`
}

// loadItemsToMap will load all the items into a map so they can be easily selected.
func loadItemsToMap() {
	ItemsMap = make(map[int]Item)
	i := Items{}
	file, err := os.Open("items.json")
	if err != nil {
		log.Fatal("couldn't open items.json file: ", err)
	}
	defer file.Close()
	data, _ := ioutil.ReadAll(file)

	err = json.Unmarshal(data, &i)
	if err != nil {
		log.Fatal("couldn't unmarshal items.json file: ", err)
	}

	for _, v := range i.Items {
		ItemsMap[v.ID] = v
	}

	return
}

func getConfiguration() (c Config) {
	dat, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal("couldn't open config.json file: ", err)
	}

	if err = json.Unmarshal(dat, &c); err != nil {
		log.Fatal("couldn't unmarshal config.json file: ", err)
	}

	return
}

func init() {
	config := getConfiguration()
	switch config.Storage {
	case "mongodb":
		mdb = MongoDBConnection{}
	case "test":
		mdb = TestDB{}
	}

	loadItemsToMap()
}

// The main function which starts the rpg
func main() {
	router := gin.Default()
	v1 := router.Group(APIBASE)
	{
		v1.GET("/", index)
		v1.POST("/create", create)
		v1.GET("/load/:name", loadCharacter)
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
