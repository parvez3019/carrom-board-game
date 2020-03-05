package command

import (
	"clean-strike/carrom"
	"clean-strike/player"
	"testing"
	"github.com/magiconair/properties/assert"
)

func Test_ShouldReturnNoErrorAndDeductScoreWhenPlayerScoreNothing(t *testing.T) {
	board := carrom.NewCarromBoard(1, 0)
	player := player.NewPlayer("Player")

	strike := NewNone()
	err := strike.Execute(player, board)

	assert.Equal(t, err, nil)
	assert.Equal(t, player.Score(), -1)
	assert.Equal(t, board.HasBlackCoins(1), true)
}

func Test_ShouldReturnNoErrorAndDecreaseScoreByOnThreeConsecutiveNoneStrikes(t *testing.T) {
	board := carrom.NewCarromBoard(1, 0)
	player := player.NewPlayer("Player")

	strike := NewNone()
	strike.Execute(player, board)
	strike.Execute(player, board)
	strike.Execute(player, board)

	assert.Equal(t, player.Score(), -4)
	assert.Equal(t, board.HasBlackCoins(1), true)
}

func Test_ShouldReturnNoErrorAndDecreaseScoreByOnSixConsecutiveNonStrikes(t *testing.T) {
	board := carrom.NewCarromBoard(1, 0)
	player := player.NewPlayer("Player")

	strike := NewNone()
	strike.Execute(player, board)
	strike.Execute(player, board)
	strike.Execute(player, board)
	strike.Execute(player, board)
	strike.Execute(player, board)
	strike.Execute(player, board)

	assert.Equal(t, player.Score(), -8)
	assert.Equal(t, board.HasBlackCoins(1), true)
}
