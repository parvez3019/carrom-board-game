package command

import (
	"clean-strike/player"
	"clean-strike/carrom"
	"errors"
)

type RedStrike struct {
}

func NewRedStrike() Command {
	return new(RedStrike)
}

func (*RedStrike) Execute(player *player.Player, board *carrom.Board) error {
	if !board.HasRedCoins(1) {
		return errors.New("not enough coins on the board")
	}
	board.PocketRedCoins(1)
	player.UpdateScore(3)
	return nil
}

func (*RedStrike) ExecuteWithCoin(player *player.Player, board *carrom.Board, coin string) error {
	return nil
}