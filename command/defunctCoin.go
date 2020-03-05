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
	return nil
}

func (*DefunctCoin) ExecuteWithCoin(player *player.Player, board *carrom.Board, coin string) error {
	if coin == carrom.RED && board.HasRedCoins(1) {
		board.PocketRedCoins(1)
	} else if coin == carrom.BLACK && board.HasBlackCoins(1) {
		board.PocketBlackCoins(1)
	} else if coin == carrom.RED || coin == carrom.BLACK {
		return errors.New("not enough coins on the board")
	} else {
		return errors.New("invalid coin")
	}
	player.UpdateScore(-2)
	return nil
}
