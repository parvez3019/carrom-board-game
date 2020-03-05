package command

import (
	"clean-strike/carrom"
	"clean-strike/player"
	"testing"
	"github.com/magiconair/properties/assert"
)

func Test_ShouldReturnNoErrorWhenPlayerScoreMultiStrike(t *testing.T) {
	board := carrom.NewCarromBoard(2, 0)
	player := player.NewPlayer("Player")

	strike := NewMultiStrike()
	err := strike.Execute(player, board)

	assert.Equal(t, err, nil)
	assert.Equal(t, 2, player.Score())
	assert.Equal(t, board.HasAllCoinsExhausted(), true)
}

func Test_ShouldReturnErrorWhenPlayerMultiStrikeButNoCoinsOnBoard(t *testing.T) {
	board := carrom.NewCarromBoard(0, 0)
	player := player.NewPlayer("Player")

	strike := NewMultiStrike()
	err := strike.Execute(player, board)

	assert.Equal(t, err.Error(), "not enough coins on the board")
	assert.Equal(t, 0, player.Score())
}

func Test_ShouldReturnErrorWhenPlayerMultiStrikeButThereIsOnlyOneCoinOnBoard(t *testing.T) {
	board := carrom.NewCarromBoard(1, 0)
	player := player.NewPlayer("Player")

	strike := NewMultiStrike()
	err := strike.Execute(player, board)

	assert.Equal(t, err.Error(), "not enough coins on the board")
	assert.Equal(t, 0, player.Score())
}

func Test_ShouldReturnNoErrorWhenPlayerScoreMultiStrikeMultipleTimes(t *testing.T) {
	board := carrom.NewCarromBoard(6, 0)
	player := player.NewPlayer("Player")

	strike := NewMultiStrike()
	strike.Execute(player, board)
	strike.Execute(player, board)
	strike.Execute(player, board)

	assert.Equal(t, 6, player.Score())
	assert.Equal(t, board.HasAllCoinsExhausted(), true)
}
