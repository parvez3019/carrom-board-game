package cleanStrike

import (
	"clean-strike/command"
	"clean-strike/carrom"
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
		STRIKE:        command.NewStrike,
		MULTISTRIKE:   command.NewMultiStrike,
		REDSTRIKE:     command.NewRedStrike,
		STRIKERSTRIKE: command.NewStrikerStrike,
		DEFUNCTCOIN:   command.NewDefunctCoin,
		NONE:          command.NewNone,
	}
	return &Game{
		Board:         board,
		commands:      commands,
	}
}
func (this *Game) SetCurrentPlayer(player *player.Player) *Game {
	this.currentPlayer = player
	return this
}

func (this *Game) Move(strike string, optionalCoin string) error {
	command := this.commands[strike]
	if command == nil {
		return errors.New("invalid command")
	}

	var err error
	if strike == DEFUNCTCOIN {
		err = command().ExecuteWithCoin(this.currentPlayer, this.Board, optionalCoin)
	} else {
		err = command().Execute(this.currentPlayer, this.Board)
	}
	if err != nil {
		return err
	}
	return nil
}

func (this *Game) CurrentPlayerName() string  {
	return this.currentPlayer.Name()
}

func (this *Game) CanContinue() bool {
	return !this.HasAllCoinsExhausted()
}

func (this *Game) Result(player1 *player.Player, player2 *player.Player) string {
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
