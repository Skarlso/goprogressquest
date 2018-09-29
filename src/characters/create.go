package characters

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/Skarlso/goprogressquest/src/responsetypes"
	"github.com/gin-gonic/gin"
)

// ItemsMap contains all the items from items.json file
var ItemsMap = make(map[int]Item, 0)

//NewCharacter The ID of a newly created character
type NewCharacter struct {
	CharacterID string `json:"id"`
}

// Create handling the creation of a new character
// curl -H "Content-Type: application/json" -X POST -d '{"name":"asdf"}' http://localhost:8989
func Create(c *gin.Context) {

	var newName struct {
		Name string `json:"name"`
	}
	ch := NewCharacter{}

	if err := c.BindJSON(&newName); err != nil {
		c.JSON(http.StatusBadRequest, responsetypes.ErrorResponse{ErrorMessage: "error while binding newName:" + err.Error()})
		return
	}

	checkSum := sha1.Sum([]byte(newName.Name))
	ch.CharacterID = fmt.Sprintf("%x", checkSum)

	char := createCharacter(ch.CharacterID, newName.Name)

	log.Println("Saving character:", char)
	err := DB.Save(char)
	if err != nil {
		c.JSON(http.StatusBadRequest, responsetypes.ErrorResponse{ErrorMessage: "error while saving character:" + err.Error()})
		return
	}
	c.JSON(http.StatusCreated, char)
}

// loadItemsToMap will load all the items into a map so they can be easily selected.
func loadItemsToMap() {
	ItemsMap = make(map[int]Item)
	i := Items{}
	file, err := os.Open("./configs/items.json")
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

func LoadCharacter(c *gin.Context) {
	charName := c.Param("name")
	var resultCharacter Character
	log.Println("Looking for character with ID:", charName)

	resultCharacter, err := DB.Load(charName)
	if err != nil {
		c.JSON(http.StatusBadRequest, responsetypes.ErrorResponse{ErrorMessage: "Error occured while loading character:" + err.Error()})
		return
	}

	c.JSON(http.StatusOK, resultCharacter)
}

func createCharacter(id, name string) Character {
	w := Item{Name: "Basic Sword", ID: 1, Armor: 0, Dmg: 10, Weight: 1, Value: 10}
	ch := Character{
		ID:   id,
		Name: name,
		Race: selectRandomRace(),
		Cast: selectRandomCast(),
		Gold: 0,
		Inventory: Inventory{
			Items:    []Item{},
			Capacity: 80,
		},
		Level:       1,
		Hp:          100,
		MaxHp:       100,
		CurrentXp:   0,
		NextLevelXp: 1000,
		//Starts out with everything at 5. Varies randomly later.
		Stats: Stats{
			Strenght:     5,
			Intelligence: 5,
			Luck:         5,
			Perception:   5,
			Agility:      5,
			Constitution: 5,
		},
		Body: Body{
			LRing:   Item{},
			RRing:   Item{},
			Armor:   Item{},
			Head:    Item{},
			Weapond: w,
			Shield:  Item{},
		},
	}

	return ch
}

func selectRandomRace() int {
	return selectRandomAffiliation("races")

}

func selectRandomCast() int {
	return selectRandomAffiliation("casts")
}

func selectRandomAffiliation(t string) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	dat, err := ioutil.ReadFile(t + ".json")
	if err != nil {
		log.Println("unable to load source file for affiliations. returning 0")
		return 0
	}
	var m map[string]interface{}
	if err = json.Unmarshal(dat, &m); err != nil {
		log.Println("unable to unmarshal. returning 0")
		return 0
	}

	log.Println(m)
	elements := m[t].([]interface{})
	elem := elements[r.Intn(len(elements)-1)]
	return int(elem.(map[string]interface{})["id"].(float64))
}
