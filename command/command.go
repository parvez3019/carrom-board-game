package command

import (
	"clean-strike/player"
	"clean-strike/carrom"
)

type Command interface {
	Execute(*player.Player, *carrom.Board) error
	ExecuteWithCoin(*player.Player, *carrom.Board, string) error
}
