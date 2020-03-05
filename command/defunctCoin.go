package command

import (
	"clean-strike/player"
	"clean-strike/carrom"
	"errors"
)

type DefunctCoin struct {
}

func NewDefunctCoin() Command {
	return new(DefunctCoin)
}

func (*DefunctCoin) Execute(player *player.Player, board *carrom.Board) error {
	if !board.HasBlackCoins(1) {
		return errors.New("not enough coins on the board")
	}
	board.DefunctCoins(1)
	player.UpdateScore(-2)
	return nil
}
