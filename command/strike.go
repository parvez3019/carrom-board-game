package command

import (
	"clean-strike/player"
	"clean-strike/carrom"
	"github.com/go-errors/errors"
)

type Strike struct {
}

func NewStrike() Command {
	return new(Strike)
}

func (*Strike) Execute(player *player.Player, board *carrom.Board) error {
	if !board.HasBlackCoins(1) {
		return errors.New("not enough coins on the board")
	}
	board.PocketBlackCoins(1)
	player.UpdateScore(1)
	return nil
}

func (*Strike) ExecuteWithCoin(player *player.Player, board *carrom.Board, coin string) error {
	return nil
}