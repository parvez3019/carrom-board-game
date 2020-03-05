package command

import (
	"clean-strike/player"
	"clean-strike/carrom"
	"errors"
)

type MultiStrike struct {
}

func NewMultiStrike() Command {
	return new(MultiStrike)
}

func (*MultiStrike) Execute(player *player.Player, board *carrom.Board) error {
	if !board.HasBlackCoins(2) {
		return errors.New("not enough coins on the board")
	}
	board.PocketBlackCoins(2)
	player.UpdateScore(2)
	return nil
}
