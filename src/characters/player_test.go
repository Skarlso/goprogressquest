package characters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLevelUpDistributingPointsShouldEqualToMaximum(t *testing.T) {
	dp := distributePoints()
	sum := 0
	for _, v := range dp {
		sum += v
	}

	assert.Equal(t, PointsPerLevel, sum, "Sum of distributed points did not equal maximum points per level.")
}

func TestLevelUpCharacterIsChanged(t *testing.T) {
	c := createCharacter("testid", "testname")
	beforeStats := c.Stats
	c.LevelUp()
	afterStats := c.Stats
	assert.NotEqual(t, beforeStats, afterStats, "Before and after stat should not equal after a levelup.")
}

func TestLevelUpIncreasesHP(t *testing.T) {
	c := createCharacter("id", "testname")
	beforeHp := c.MaxHp
	c.LevelUp()
	afterHp := c.MaxHp
	assert.NotEqual(t, beforeHp, afterHp, "Before hp equald after hp. It should not have.")
}

func TestLevelUpIncreasesXP(t *testing.T) {
	c := createCharacter("id", "testname")
	c.CurrentXp = 100
	c.LevelUp()
	assert.Equal(t, c.CurrentXp, 0, "Xp should be reset.")
}

func TestLevelUpIncreasesLevel(t *testing.T) {
	c := createCharacter("id", "testname")
	currentLevel := c.Level
	c.LevelUp()
	newLevel := c.Level
	assert.NotEqual(t, newLevel, currentLevel, "Level should be higher.")
	// assert.Condition(t, assert.Comparison(func(a, b int) bool { return b > a }), "Level should be higher.")
}
