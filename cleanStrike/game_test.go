package cleanStrike

import (
	"clean-strike/carrom"
	"clean-strike/player"
	"testing"
	"github.com/magiconair/properties/assert"
)

func setUpGame(blackCoins int, redCoins int) (*Game, *player.Player, *player.Player,) {
	board := carrom.NewCarromBoard(blackCoins, redCoins)
	player1 := player.NewPlayer("Player1")
	player2 := player.NewPlayer("Player2")
	game := NewGame(board)

	return game, player1, player2
}

func TestCleanStrike_ShouldReturnGameAsDrawWithScore2and0As0_WhenPlayer1StrikeTwiceInRow(t *testing.T) {
	game, player1, player2 := setUpGame(2, 0)

	game.SetCurrentPlayer(player1)
	game.Move(STRIKE, "")
	game.Move(STRIKE, "")

	assert.Equal(t, player1.Score(), 2)
	assert.Equal(t, player2.Score(), 0)
	assert.Equal(t, game.Result(player1, player2), "Game is draw")
}

func TestCleanStrike_ShouldReturnAsDrawEvenWith3PlusLeadForOnePlaceIfLacksByMin5Constraint_WhenPlayer1score1AndPlayer2Score4(t *testing.T) {
	game, player1, player2 := setUpGame(9, 1)

	game.SetCurrentPlayer(player1)
	game.Move(STRIKE, "") // P1: 1

	game.SetCurrentPlayer(player2)
	game.Move(STRIKE, "") // P2: 1

	game.SetCurrentPlayer(player1)
	game.Move(MULTISTRIKE, "") // P1: 3

	game.SetCurrentPlayer(player2)
	game.Move(REDSTRIKE, "")   // P2: 4

	game.SetCurrentPlayer(player1)
	game.Move(STRIKERSTRIKE, "")

	game.SetCurrentPlayer(player2) // P1: 2
	game.Move(MULTISTRIKE, "")   // P2: 6

	game.SetCurrentPlayer(player1)
	game.Move(NONE, "")                  // P1: 1

	game.SetCurrentPlayer(player2)
	game.Move(DEFUNCTCOIN, carrom.BLACK) // P2: 4

	assert.Equal(t, player1.Score(), 1)
	assert.Equal(t, player2.Score(), 4)
	assert.Equal(t, game.Result(player1, player2), "Game is draw")
}


func TestCleanStrike_ShouldReturnAsDrawPlayer2AsWinner_WhenPlayer1score1AndPlayer2Score5(t *testing.T) {
	game, player1, player2 := setUpGame(9, 1)

	game.SetCurrentPlayer(player1)
	game.Move(STRIKE, "") // P1: 1

	game.SetCurrentPlayer(player2)
	game.Move(STRIKE, "") // P2: 1

	game.SetCurrentPlayer(player1)
	game.Move(MULTISTRIKE, "") // P1: 3

	game.SetCurrentPlayer(player2)
	game.Move(REDSTRIKE, "")   // P2: 4

	game.SetCurrentPlayer(player1)
	game.Move(STRIKERSTRIKE, "") // P1: 2

	game.SetCurrentPlayer(player2)
	game.Move(MULTISTRIKE, "")  // P2: 6

	game.SetCurrentPlayer(player1)
	game.Move(NONE, "")      // P1: 1

	game.SetCurrentPlayer(player2)
	game.Move(DEFUNCTCOIN, carrom.BLACK) // P2: 4

	game.SetCurrentPlayer(player1)
	game.Move(STRIKE, "") // P1: 2

	game.SetCurrentPlayer(player2)
	game.Move(STRIKE, "") // P2: 5

	assert.Equal(t, player1.Score(), 2)
	assert.Equal(t, player2.Score(), 5)
	assert.Equal(t, game.Result(player1, player2), "Player2 won the game. Final Score: 5 : 2")
}

func TestCleanStrike_ShouldReturnWinnerAsPlayer1_WhenPlayer1score5AndPlayer2Score0(t *testing.T) {
	game, player1, player2 := setUpGame(9, 1)

	player1.UpdateScore(5)
	player2.UpdateScore(0)

	assert.Equal(t, game.Result(player1, player2), "Player1 won the game. Final Score: 5 : 0")
}

func TestCleanStrike_ShouldReturnDraw_WhenPlayer1score3AndPlayer2Score5(t *testing.T) {
	game, player1, player2 := setUpGame(9, 1)

	player1.UpdateScore(3)
	player2.UpdateScore(5)

	assert.Equal(t, game.Result(player1, player2), "Game is draw")
}
