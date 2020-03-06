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
	game := NewCleanStrike(board)

	return game, player1, player2
}

func TestCleanStrike_ShouldReturnGameAsDrawWithScore2and0As0_WhenPlayer1StrikeTwiceInRow(t *testing.T) {
	game, player1, player2 := setUpGame(2, 0)

	game.Play(player1, STRIKE, "")
	game.Play(player1, STRIKE, "")

	assert.Equal(t, player1.Score(), 2)
	assert.Equal(t, player2.Score(), 0)
	assert.Equal(t, game.Result(player1, player2), "Game is draw")
}

func TestCleanStrike_ShouldReturnAsDrawEvenWith3PlusLeadForOnePlaceIfLacksByMin5Constraint_WhenPlayer1score1AndPlayer2Score4(t *testing.T) {
	game, player1, player2 := setUpGame(9, 1)

	game.Play(player1, STRIKE, "") // P1: 1
	game.Play(player2, STRIKE, "") // P2: 1

	game.Play(player1, MULTISTRIKE, "") // P1: 3
	game.Play(player2, REDSTRIKE, "")   // P2: 4

	game.Play(player1, STRIKERSTRIKE, "") // P1: 2
	game.Play(player2, MULTISTRIKE, "")   // P2: 6

	game.Play(player1, NONE, "")                  // P1: 1
	game.Play(player2, DEFUNCTCOIN, carrom.BLACK) // P2: 4

	assert.Equal(t, player1.Score(), 1)
	assert.Equal(t, player2.Score(), 4)
	assert.Equal(t, game.Result(player1, player2), "Game is draw")
}


func TestCleanStrike_ShouldReturnAsDrawPlayer2AsWinner_WhenPlayer1score1AndPlayer2Score5(t *testing.T) {
	game, player1, player2 := setUpGame(9, 1)

	game.Play(player1, STRIKE, "") // P1: 1
	game.Play(player2, STRIKE, "") // P2: 1

	game.Play(player1, MULTISTRIKE, "") // P1: 3
	game.Play(player2, REDSTRIKE, "")   // P2: 4

	game.Play(player1, STRIKERSTRIKE, "") // P1: 2
	game.Play(player2, MULTISTRIKE, "")   // P2: 6

	game.Play(player1, NONE, "")                  // P1: 1
	game.Play(player2, DEFUNCTCOIN, carrom.BLACK) // P2: 4

	game.Play(player1, STRIKE, "") // P1: 2
	game.Play(player2, STRIKE, "") // P2: 5

	assert.Equal(t, player1.Score(), 2)
	assert.Equal(t, player2.Score(), 5)
	assert.Equal(t, game.Result(player1, player2), "Player2 won the game. Final Score: 5-2")
}

func TestCleanStrike_ShouldReturnWinnerAsPlayer1_WhenPlayer1score5AndPlayer2Score0(t *testing.T) {
	game, player1, player2 := setUpGame(9, 1)

	player1.UpdateScore(5)
	player2.UpdateScore(0)

	assert.Equal(t, game.Result(player1, player2), "Player1 won the game. Final Score: 5-0")
}

func TestCleanStrike_ShouldReturnDraw_WhenPlayer1score3AndPlayer2Score5(t *testing.T) {
	game, player1, player2 := setUpGame(9, 1)

	player1.UpdateScore(3)
	player2.UpdateScore(5)

	assert.Equal(t, game.Result(player1, player2), "Game is draw")
}
