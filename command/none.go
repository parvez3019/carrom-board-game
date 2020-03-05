package command

import (
	"clean-strike/player"
	"clean-strike/carrom"
)

type None struct {
}

func NewNone() Command {
	return new(None)
}

func (*None) Execute(player *player.Player, board *carrom.Board) error {
	player.Foul()
	return nil
}

func (*None) ExecuteWithCoin(player *player.Player, board *carrom.Board, coin string) error {
	return nil
}