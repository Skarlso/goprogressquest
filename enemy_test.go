package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	mdb = TestDB{}
}

func TestEnemyIsInitializedFromJSONFile(t *testing.T) {
	c := createCharacter("id", "test")
	e := SpawnEnemy(c)

	assert.NotEmpty(t, e)
}
