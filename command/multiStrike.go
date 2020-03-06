package command

import (
	"clean-strike/player"
	"clean-strike/carrom"
	"errors"
	"clean-strike/constants"
)

type MultiStrike struct {
}

func NewMultiStrike() Command {
	return new(MultiStrike)
}

func (*MultiStrike) Execute(player *player.Player, board *carrom.Board) error {
	if !board.HasBlackCoins(2) {
		return errors.New(constants.NotEnoughCoinsError)
	}
	board.RemoveNBlackCoins(2)
	player.UpdateScore(2)
	return nil
}

func (*MultiStrike) ExecuteWithCoin(player *player.Player, board *carrom.Board, coin string) error {
	return nil
}
