package main

import (
	"clean-strike/player"
	"fmt"
	"clean-strike/carrom"
	"clean-strike/cleanStrike"
)

func main() {
	board := carrom.NewCarromBoard(9, 1)
	player1 := player.NewPlayer("Player 1")
	player2 := player.NewPlayer("Player 2")

	cs := cleanStrike.NewCleanStrike(board)
	gameRunner := cleanStrike.NewGameRunner(cs)

	result := gameRunner.Play(player1, player2)

	fmt.Println(result)
}
