package command

import (
	"clean-strike/carrom"
	"clean-strike/player"
	"testing"
	"github.com/magiconair/properties/assert"
)

func Test_ShouldReturnNoErrorWhenPlayerScoreRedStrike(t *testing.T) {
	board := carrom.NewCarromBoard(1, 1)
	player := player.NewPlayer("Player")

	strike := NewRedStrike()
	err := strike.Execute(player, board)

	assert.Equal(t, err, nil)
	assert.Equal(t, 3, player.Score())
	assert.Equal(t, board.HasBlackCoins(1), true)
	assert.Equal(t, board.HasRedCoins(1), false)
}

func Test_ShouldReturnErrorWhenPlayerRedStrikeButNoCoinsOnBoard(t *testing.T) {
	board := carrom.NewCarromBoard(1, 0)
	player := player.NewPlayer("Player")

	strike := NewRedStrike()
	err := strike.Execute(player, board)

	assert.Equal(t, err.Error(), "not enough coins on the board")
	assert.Equal(t, 0, player.Score())
	assert.Equal(t, board.HasBlackCoins(1), true)

}

func Test_ShouldReturnNoErrorWhenPlayerScoreRedStrikeMultipleTimes(t *testing.T) {
	board := carrom.NewCarromBoard(5, 5)
	player := player.NewPlayer("Player")

	strike := NewRedStrike()
	strike.Execute(player, board)
	strike.Execute(player, board)
	strike.Execute(player, board)

	assert.Equal(t, 9, player.Score())
	assert.Equal(t, board.HasRedCoins(3), false)
}
