package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	mdb = TestDB{}
}

func TestEnemyIsInitializedFromJSONFile(t *testing.T) {
	c := createCharacter("id", "test")
	e := SpawnEnemy(c)

	fmt.Println("Enemy:", e)
	assert.NotEmpty(t, e)
}
