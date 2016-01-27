Go Progress Quest
=================

[![Build Status](https://travis-ci.org/Skarlso/goprogressquest.svg?branch=master)](https://travis-ci.org/Skarlso/goprogressquest) [![Coverage Status](https://coveralls.io/repos/github/Skarlso/goprogressquest/badge.svg?branch=master)](https://coveralls.io/github/Skarlso/goprogressquest?branch=master)

Coverage Report (Until I get Coveralls working...)
--------------------------------------------------

```bash
go test -coverprofile=cover.out
go tool cover -html=cover.out
```

Cover report - 27.01.2016
-------------------------



This will be a Go Implementation of the famous type of RPG, called:
https://en.wikipedia.org/wiki/Progress_Quest

This will be an API which can be consumed by any client in a number of ways.

API Version is 1
----------------

/api/1/*

The following end-points are available:

Creational
----------

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
/create
```
```json
POST:
{
    "name":"MyAwesomeCharacterName"
}
```
Return:
```json
{
    "id":"103b922810b1fac97da1bad872618477"
}
```

```bash
# Load a character by ID since names are not unique at the moment
/load/{ID}
/load/3da541559918a808c2402bba5012f6c60b27661c
```

Return:
```json
{
    "Inventory":{"Items":[]},"Name":"MyAwesomeCharacterName","Stats":{"Str":0,"Agi":0,"In":0,"Per":0,"Chr":0,"Lck":0},"ID":"3da541559918a808c2402bba5012f6c60b27661c","Gold":0
}
```

Adventuring related
-------------------

```bash
# Start adventuring
/start
```

```json
POST:
{
    "name":"MyAwesomeCharacterName"
}
```
Return:
```json
{
    "message":"Started adventuring for character: MyAwesomeCharacterName"
}
```

```bash
# Stop adventuring
/stop
```
```json
POST:
{
    "name":"MyAwesomeCharacterName"
}
```
Return:
```json
{
    "message":"Stopped adventuring for character: MyAwesomeCharacterName"
}
```

Running it
----------

```bash
go build
```

Currently the project is simple enough so that no Makefile is needed for this process.
