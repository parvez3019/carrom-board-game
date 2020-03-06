package cleanStrike

import (
	"clean-strike/command"
	"clean-strike/carrom"
	"github.com/go-errors/errors"
	"clean-strike/player"
	"fmt"
)

type CleanStrike struct {
	*carrom.Board
	commands map[string]func() command.Command
}

func NewCleanStrike(board *carrom.Board) *CleanStrike {
	commands := map[string]func() command.Command{
		STRIKE:        command.NewStrike,
		MULTISTRIKE:   command.NewMultiStrike,
		REDSTRIKE:     command.NewRedStrike,
		STRIKERSTRIKE: command.NewStrikerStrike,
		DEFUNCTCOIN:   command.NewDefunctCoin,
		NONE:          command.NewNone,
	}
	return &CleanStrike{
		Board:    board,
		commands: commands,
	}
}

func (this *CleanStrike) Move(player *player.Player, strike string, optionalCoin string) error {
	command := this.commands[strike]
	if command == nil {
		return errors.New("invalid command")
	}

	var err error
	if strike == DEFUNCTCOIN {
		err = command().ExecuteWithCoin(player, this.Board, optionalCoin)
	} else {
		err = command().Execute(player, this.Board)
	}
	if err != nil {
		return err
	}
	return nil
}

func (this *CleanStrike) CanContinue() bool  {
	return !this.HasAllCoinsExhausted()
}

func (this *CleanStrike) Result(player1 *player.Player, player2 *player.Player) string {
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
