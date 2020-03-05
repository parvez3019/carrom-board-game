package main

import (
	"clean-strike/command"
	"clean-strike/carrom"
	"github.com/go-errors/errors"
	"clean-strike/player"
	"strconv"
)

type CleanStrike struct {
	// players
	// carrom-board
	// Scoreboard
	// Commands
	// currentTurn
	*carrom.Board
	players  []*player.Player
	commands map[string]func() command.Command
}

func NewCleanStrike(players []*player.Player, board *carrom.Board) *CleanStrike {
	commands := map[string]func() command.Command{
		"strike": command.NewStrike,
		"multi-strike":   command.NewMultiStrike,
		"red-strike":     command.NewRedStrike,
		"striker-strike": command.NewStrikerStrike,
		"defunct-coin":   command.NewDefunctCoin,
		"none":           command.NewNone,
	}
	return &CleanStrike{
		Board:    board,
		players:  players,
		commands: commands,
	}
}

func (this *CleanStrike) Play(player *player.Player, strike string) error {
	command := this.commands[strike]
	if command == nil {
		return errors.New("Invalid Command")
	}

	err := command().Execute(player, this.Board)
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
