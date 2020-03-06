package command

import (
	"clean-strike/player"
	"clean-strike/carrom"
	"errors"
	"clean-strike/cleanStrike"
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
		board.RemoveNRedCoins(1)
	} else if coin == carrom.BLACK && board.HasBlackCoins(1) {
		board.RemoveNBlackCoins(1)
	} else if coin == carrom.RED || coin == carrom.BLACK {
		return errors.New(cleanStrike.NotEnoughCoinsError)
	} else {
		return errors.New(cleanStrike.InvalidCoinError)
	}
	player.UpdateScore(-2)
	return nil
}
