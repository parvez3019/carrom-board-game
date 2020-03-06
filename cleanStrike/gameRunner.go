package cleanStrike

import (
	"clean-strike/player"
	"fmt"
	"strings"
)

type GameRunner struct {
	*Game
}

func NewGameRunner(strike *Game) *GameRunner {
	return &GameRunner{Game: strike}
}

func (g *GameRunner) Play(player1 *player.Player, player2 *player.Player) string {
	var err error
	var currentPlayer *player.Player
	for g.CanContinue() {
		if err == nil {
			currentPlayer = g.getNextPlayer(currentPlayer, player1, player2)
		}
		g.SetCurrentPlayer(currentPlayer)
		err = g.makeMove()
		if err != nil && err.Error() == NotEnoughCoinsError {
			fmt.Println(NotEnoughCoinsErrorMessage)
		} else if err != nil {
			fmt.Println(InvalidCommandErrorMessage)
		}
	}
	return g.Result(player1, player2)
}

func (g *GameRunner) makeMove() error {
	commandChooseByPlayer, optionalCoinInfo, err := g.getInputFromUser()
	if err != nil {
		return err
	}
	return g.Move(inputCommandMap[commandChooseByPlayer], optionalCoinInfo)
}

func (g *GameRunner) getInputFromUser() (string, string, error) {
	g.printDisplayInputMessage()
	return g.getInputCommandAndOptionCoinColor()
}

func (g *GameRunner) getInputCommandAndOptionCoinColor() (string, string, error) {
	var inputCommand string
	var optionalCoin string
	_, err := fmt.Scanln(&inputCommand)
	if err != nil {
		return inputCommand, optionalCoin, err
	}
	if inputCommandMap[inputCommand] == DEFUNCTCOIN {
		_, err := fmt.Scanln(&optionalCoin)
		if err != nil {
			return inputCommand, optionalCoin, err
		}
		optionalCoin = strings.Trim(strings.ToLower(optionalCoin), " ")
	}
	return inputCommand, optionalCoin, err
}

func (g *GameRunner) printDisplayInputMessage() {
	fmt.Println(fmt.Sprintf("%s : Choose an outcome from the list below", g.CurrentPlayerName()))
	fmt.Println("1. Strike")
	fmt.Println("2. Multistrike")
	fmt.Println("3. Red strike")
	fmt.Println("4. Striker strike")
	fmt.Println("5. Defunct coin")
	fmt.Println("6. None")
}

func (g *GameRunner) getNextPlayer(current *player.Player, p1 *player.Player, p2 *player.Player) *player.Player {
	if current == nil {
		return p1
	}
	if current.Id() == p1.Id() {
		return p2
	}
	return p1
}

var inputCommandMap = map[string]string{
	"1": STRIKE,
	"2": MULTISTRIKE,
	"3": REDSTRIKE,
	"4": STRIKERSTRIKE,
	"5": DEFUNCTCOIN,
	"6": NONE,
}
