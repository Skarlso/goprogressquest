package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThatMonsterLevelCannotGoBelowZero(t *testing.T) {
	ch := createCharacter("id", "test")
	ch.Level = -10
	monster := ch.spawnEnemy()

	assert.Equal(t, monster.Level, 0, "Monster level did not match 0.")
}
