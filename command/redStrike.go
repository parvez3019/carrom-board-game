package command

import (
	"clean-strike/player"
	"clean-strike/carrom"
	"errors"
	"clean-strike/cleanStrike"
)

type RedStrike struct {
}

func NewRedStrike() Command {
	return new(RedStrike)
}

func (*RedStrike) Execute(player *player.Player, board *carrom.Board) error {
	if !board.HasRedCoins(1) {
		return errors.New(cleanStrike.NotEnoughCoinsError)
	}
	board.RemoveNRedCoins(1)
	player.UpdateScore(3)
	return nil
}

func (*RedStrike) ExecuteWithCoin(player *player.Player, board *carrom.Board, coin string) error {
	return nil
}