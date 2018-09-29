package characters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnemyIsInitializedFromJSONFile(t *testing.T) {
	c := createCharacter("id", "test")
	e := SpawnEnemy(c)

	assert.NotEmpty(t, e)
}
