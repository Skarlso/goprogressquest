Changes to Go ProgressQuest
===========================


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
