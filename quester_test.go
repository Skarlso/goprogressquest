package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThatMonsterLevelCannotGoBelowZero(t *testing.T) {
	ch := createCharacter("id", "test")
	ch.Level = -10
	monster := ch.spawnEnemy()
	assert.Condition(t, func() bool { return monster.Level >= 0 }, "Monster level was below 0.")
	// assert.Equal(t, monster.Level, 0, "Monster level did not match 0.")
}
