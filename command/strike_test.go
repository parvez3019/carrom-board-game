package command

import (
	"clean-strike/carrom"
	"clean-strike/player"
	"testing"
	"github.com/magiconair/properties/assert"
)

func Test_ShouldReturnNoErrorWhenPlayerScoreStrike(t *testing.T) {
	board := carrom.NewCarromBoard(1, 0)
	player := player.NewPlayer("Player")

	strike := NewStrike()
	err := strike.Execute(player, board)

	assert.Equal(t, err, nil)
	assert.Equal(t, 1, player.Score())
	assert.Equal(t, board.HasBlackCoins(1), false)
}

func Test_ShouldReturnErrorWhenPlayerStrikeButNoCoinsOnBoard(t *testing.T) {
	board := carrom.NewCarromBoard(0, 0)
	player := player.NewPlayer("Player")

	strike := NewStrike()
	err := strike.Execute(player, board)

	assert.Equal(t, err.Error(), "not enough coins on the board")
	assert.Equal(t, 0, player.Score())
}

func Test_ShouldReturnNoErrorWhenPlayerScoreStrikeMultipleTimes(t *testing.T) {
	board := carrom.NewCarromBoard(5, 0)
	player := player.NewPlayer("Player")

	strike := NewStrike()
	strike.Execute(player, board)
	strike.Execute(player, board)
	strike.Execute(player, board)

	assert.Equal(t, 3, player.Score())
	assert.Equal(t, board.HasBlackCoins(3), false)
}
