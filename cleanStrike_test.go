package main

import (
	"clean-strike/carrom"
	"clean-strike/player"
	"testing"
	"github.com/magiconair/properties/assert"
)

func setUpGame(blackCoins int, redCoins int) (*CleanStrike, *player.Player, *player.Player,) {
	board := carrom.NewCarromBoard(blackCoins, redCoins)
	player1 := player.NewPlayer("Player1")
	player2 := player.NewPlayer("Player2")
	game := NewCleanStrike([]*player.Player{player1, player2}, board)

	return game, player1, player2
}

func TestCleanStrike_ShouldReturnResultAs1dash0WhenPlayer1Strike(t *testing.T) {
	game, player1, _ := setUpGame(2, 0)

	game.Play(player1, "strike")

	assert.Equal(t, game.Result(), "1-0")

}
//
//func TestCleanStrike_ShouldReturnResultAs1dash1WhenPlayer1and2BothStrike(t *testing.T) {
//	game, player1, player2 := setUpGame(2, 0)
//
//	game.Play(player1, strike)
//	game.Play(player2, strike)
//
//	assert.Equal(t, game.Result(), "1-1")
//
//}
//
//func TestCleanStrike_ShouldReturnErrorWhenPlayerStrikeButNoCoinsOnBoard(t *testing.T) {
//	game, player1, _ := setUpGame(0, 0)
//
//	err := game.Play(player1, strike)
//
//	assert.Equal(t, err.Error(), "not enough coins on the board")
//
//}
//
//func TestCleanStrike_ShouldReturnResultWhenAllCoinsHasExausted(t *testing.T) {
//	game, player1, _ := setUpGame(1, 0)
//
//	err := game.Play(player1, strike)
//
//	assert.Equal(t, err.Error(), "not enough coins on the board")
//
//}