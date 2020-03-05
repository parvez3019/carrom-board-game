package command

import (
	"clean-strike/carrom"
	"clean-strike/player"
	"testing"
	"github.com/magiconair/properties/assert"
)

func Test_ShouldReturnNoErrorAndDeductPointsWhenPlayerDefunctABlackCoin(t *testing.T) {
	board := carrom.NewCarromBoard(2, 2)
	player := player.NewPlayer("Player")

	strike := NewDefunctCoin()
	err := strike.ExecuteWithCoin(player, board, carrom.BLACK)

	assert.Equal(t, err, nil)
	assert.Equal(t, -2, player.Score())
	assert.Equal(t, board.HasBlackCoins(2), false)
	assert.Equal(t, board.HasBlackCoins(1), true)
	assert.Equal(t, board.HasRedCoins(2), true)
}

func Test_ShouldReturnNoErrorAndDeductPointsWhenPlayerDefunctARedCoin(t *testing.T) {
	board := carrom.NewCarromBoard(2, 2)
	player := player.NewPlayer("Player")

	strike := NewDefunctCoin()
	err := strike.ExecuteWithCoin(player, board, carrom.RED)

	assert.Equal(t, err, nil)
	assert.Equal(t, -2, player.Score())
	assert.Equal(t, board.HasRedCoins(2), false)
	assert.Equal(t, board.HasRedCoins(1), true)
	assert.Equal(t, board.HasBlackCoins(2), true)
}

func Test_ShouldReturnNoErrorAndDeduct4PointsWhenPlayerDefunctTwoRedCoin(t *testing.T) {
	board := carrom.NewCarromBoard(2, 2)
	player := player.NewPlayer("Player")

	strike := NewDefunctCoin()
	strike.ExecuteWithCoin(player, board, carrom.RED)
	strike.ExecuteWithCoin(player, board, carrom.RED)

	assert.Equal(t, -4, player.Score())
	assert.Equal(t, board.HasRedCoins(1), false)
	assert.Equal(t, board.HasBlackCoins(2), true)
}

func Test_ShouldReturnErrorWhenPlayerDefunctABlackCoinWhenNoBlackCoinsPresentInCarrom(t *testing.T) {
	board := carrom.NewCarromBoard(0, 0)
	player := player.NewPlayer("Player")

	strike := NewDefunctCoin()
	err := strike.ExecuteWithCoin(player, board, carrom.BLACK)

	assert.Equal(t, err.Error(), "not enough coins on the board")
	assert.Equal(t, 0, player.Score())
}

func Test_ShouldReturnErrorWhenPlayerDefunctARedCoinWhenNoRedCoinsPresentInCarrom(t *testing.T) {
	board := carrom.NewCarromBoard(0, 0)
	player := player.NewPlayer("Player")

	strike := NewDefunctCoin()
	err := strike.ExecuteWithCoin(player, board, carrom.RED)

	assert.Equal(t, err.Error(), "not enough coins on the board")
	assert.Equal(t, 0, player.Score())
}

func Test_ShouldReturnErrorWhenPlayerDefunctAnNonExistingColorCoin(t *testing.T) {
	board := carrom.NewCarromBoard(0, 0)
	player := player.NewPlayer("Player")

	strike := NewDefunctCoin()
	err := strike.ExecuteWithCoin(player, board, "Yellow")

	assert.Equal(t, err.Error(), "invalid coin")
	assert.Equal(t, 0, player.Score())
}
