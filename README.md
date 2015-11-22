Go Progress Quest
=================
This will be a Go Implementation of the famous type of RPG, called:
https://en.wikipedia.org/wiki/Progress_Quest

This will be an API which can be consumed by any client in a number of ways.

API Version is 1
----------------

/api/1/*

The following end-points are available:

```
/
```
Currently returns:
```json
{
    "message":"Welcome!"
}
```


```bash
# Character names don't have to be unique
/register
```
```json
POST:
{
    "character":{
        "name":"MyAwesomeCharacterName"
    }
}
```
Return: (Something like this)
```json
{
    "characterId":"103b922810b1fac97da1bad872618477"
}
```

Running it
----------

To run it, you have to install the app by typeing ```go install``` from the main directory. The switch to your go installation's bin directory and run the created goprogressquest file. That should start the server. If you are developing and don't want to swith around, ```go run``` needs all the files as parameter.

```go run gorpgapi.go middleware.go response_types.go```
