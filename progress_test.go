package main

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