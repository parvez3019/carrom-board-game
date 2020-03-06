package cleanStrike

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
	score1 := player1.Score()
	score2 := player2.Score()

	const MinimumLeadRequired = 3
	if score1 >= score2+MinimumLeadRequired && score1 >= 5 {
		return fmt.Sprintf(GameResultStringFormatter, player1.Name(), score1, score2)
	} else if score2 >= score1+MinimumLeadRequired && score2 >= 5 {
		return fmt.Sprintf(GameResultStringFormatter, player2.Name(), score2, score1)
	}
	return GameDrawString
}

const GameDrawString = "Game is draw"
const GameResultStringFormatter = "%s won the game. Final Score: %d : %d"
