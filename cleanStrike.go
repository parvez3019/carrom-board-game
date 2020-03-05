package main

import (
	"clean-strike/command"
	"clean-strike/carrom"
	"github.com/go-errors/errors"
	"clean-strike/player"
	"strconv"
)

type CleanStrike struct {
	*carrom.Board
	players  []*player.Player
	commands map[string]func() command.Command
}

func NewCleanStrike(players []*player.Player, board *carrom.Board) *CleanStrike {
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
		players:  players,
		commands: commands,
	}
}

func (this *CleanStrike) Play(player *player.Player, strike string, coin string) error {
	command := this.commands[strike]
	if command == nil {
		return errors.New("invalid command")
	}

	var err error
	if strike == DEFUNCTCOIN {
		err = command().ExecuteWithCoin(player, this.Board, coin)
	} else {
		err = command().Execute(player, this.Board)
	}
	if err != nil {
		return err
	}
	return nil
}

func (this *CleanStrike) Result() string {
	result := ""
	for _, player := range this.players {
		result = result + strconv.Itoa(player.Score()) + "-"
	}
	return result[: len(result)-1]
}
