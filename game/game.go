package game

import (
	"clean-strike/command"
	"clean-strike/carrom"
	"clean-strike/constants"
	"github.com/go-errors/errors"
	"clean-strike/player"
	"fmt"
)

type Game struct {
	*carrom.Board
	commands      map[string]func() command.Command
	currentPlayer *player.Player
}

func NewGame(board *carrom.Board) *Game {
	commands := map[string]func() command.Command{
		constants.Strike:        command.NewStrike,
		constants.MultiStrike:   command.NewMultiStrike,
		constants.RedStrike:     command.NewRedStrike,
		constants.StrikerStrike: command.NewStrikerStrike,
		constants.DefunctCoin:   command.NewDefunctCoin,
		constants.None:          command.NewNone,
	}
	return &Game{
		Board:    board,
		commands: commands,
	}
}
func (g *Game) SetCurrentPlayer(player *player.Player) *Game {
	g.currentPlayer = player
	return g
}

func (g *Game) Move(strike string, optionalCoin string) error {
	command := g.commands[strike]
	if command == nil {
		return errors.New("invalid command")
	}

	var err error
	if strike == constants.DefunctCoin {
		err = command().ExecuteWithCoin(g.currentPlayer, g.Board, optionalCoin)
	} else {
		err = command().Execute(g.currentPlayer, g.Board)
	}
	if err != nil {
		return err
	}
	return nil
}

func (g *Game) CurrentPlayerName() string {
	return g.currentPlayer.Name()
}

func (g *Game) CanContinue() bool {
	return !g.HasAllCoinsExhausted()
}

func (g *Game) Result(player1 *player.Player, player2 *player.Player) string {
	winner, loser,  draw := findWinner(player1, player2)
	if draw {
		return DrawResultString
	}
	return fmt.Sprintf(ResultStringFormatter, winner.Name() , winner.Score(), loser.Score())
}

func findWinner(player1 *player.Player, player2 *player.Player) (*player.Player, *player.Player, bool) {
	const MinimumLeadRequired = 3
	const MinPointsRequiredToWin = 5
	if player1.Score() >= player2.Score() + MinimumLeadRequired && player1.Score() >= MinPointsRequiredToWin {
		return player1, player2, false
	}
	if player2.Score() >= player1.Score() + MinimumLeadRequired && player2.Score() >= MinPointsRequiredToWin {
		return player2, player1, false
	}
	return nil, nil, true
}

const DrawResultString = "Game is draw"
const ResultStringFormatter = "%s won the game. Final Score: %d : %d"
