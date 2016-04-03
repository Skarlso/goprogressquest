Changes to Go ProgressQuest
===========================

v0.2.0
------

* Basic Fighting implemented
* Encountering an enemy implemented
* Colored log output for better visuals
* Selling items is the inventory is full (Character reached maximum bearing)
* Multiple characters can adventure at the same time

Future plans:
* Buying potion
* Buying armor / weapons
* Unique name for created characters. Currently there is no check on uniqueness
* Implementing the three possible encounters: Discovery, Encounter, Neutral
* Improve the tests


v0.1.1
------

* Starting to abstract out the DB in order to test it better
* Created a TestDB which will return true values for testing purposes
* Created a config which will load the appropriate DB
* Added loading config within init()


v0.1.0
------

* Major upgrade to use [Gin](https://github.com/gin-gonic/gin).
* Thrown out middleware since Gin prodived very nice logging and panic handling.
* Thrown out most of the error handling stuff, since Gin provides that as well.
* Very nice API version grouping thanks to Gin.
* Binding of responses to JSON and from JSON is a breeze now.
* API Listing is enabled with Gin so it's easy what API endpoint is bound to what.
* More to come...


v0.0.1
------

 * Added Licensing information and basic workings and mechanics of starting and stopping an adventure.
 * Added this change log.
