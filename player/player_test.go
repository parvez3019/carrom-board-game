package player

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_ShouldReturn2WhenPlayerUpdateTo2(t *testing.T) {
	player := NewPlayer("Player")

	player.UpdateScore(2)

	assert.Equal(t, 2, player.Score())
}

func Test_ShouldReturnMinus1WhenPlayerFoul(t *testing.T) {
	player := NewPlayer("Player")

	player.Foul()

	assert.Equal(t, -1, player.Score())
}

func Test_ShouldReturnMinus4WhenPlayerFoulForThreeConsecutiveTimes(t *testing.T) {
	player := NewPlayer("Player")

	player.Foul()
	player.Foul()
	player.Foul()

	assert.Equal(t, -4, player.Score())
}

func Test_ShouldReturnMinus8WhenPlayerFoulForSixConsecutiveTimes(t *testing.T) {
	player := NewPlayer("Player")

	player.Foul()
	player.Foul()
	player.Foul()
	player.Foul()
	player.Foul()
	player.Foul()

	assert.Equal(t, -8, player.Score())
}
