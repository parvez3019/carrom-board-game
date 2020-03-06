package command

import (
	"clean-strike/player"
	"clean-strike/carrom"
	"github.com/go-errors/errors"
	"clean-strike/cleanStrike"
)

type Strike struct {
}

func NewStrike() Command {
	return new(Strike)
}

func (*Strike) Execute(player *player.Player, board *carrom.Board) error {
	if !board.HasBlackCoins(1) {
		return errors.New(cleanStrike.NotEnoughCoinsError)
	}
	board.RemoveNBlackCoins(1)
	player.UpdateScore(1)
	return nil
}

func (*Strike) ExecuteWithCoin(player *player.Player, board *carrom.Board, coin string) error {
	return nil
}