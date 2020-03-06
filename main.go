package main

import (
	"clean-strike/player"
	"fmt"
	"clean-strike/carrom"
	"clean-strike/game"
)

func main() {
	board := carrom.NewCarromBoard(9, 1)
	player1 := player.NewPlayer("Player 1")
	player2 := player.NewPlayer("Player 2")

	cs := game.NewGame(board)
	gameRunner := game.NewGameRunner(cs)

	result := gameRunner.Play(player1, player2)

	fmt.Println(result)
}
