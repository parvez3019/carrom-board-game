package command

import (
	"clean-strike/carrom"
	"clean-strike/player"
)

type StrikerStrike struct {
}

func NewStrikerStrike() Command {
	return new(StrikerStrike)
}

func (*StrikerStrike) Execute(player *player.Player, board *carrom.Board) error {
	player.Foul()
	return nil
}

func (*StrikerStrike) ExecuteWithCoin(player *player.Player, board *carrom.Board, coin string) error {
	return nil
}