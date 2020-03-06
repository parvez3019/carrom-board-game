package cleanStrike

import (
	"clean-strike/player"
	"fmt"
	"strings"
)

type GameRunner struct {
	*CleanStrike
}

func NewGameRunner(strike *CleanStrike) *GameRunner {
	return &GameRunner{CleanStrike: strike}
}

func (this *GameRunner) Play(player1 *player.Player, player2 *player.Player) string {
	var err error
	var currentPlayer *player.Player
	for this.CanContinue() {
		if err == nil {
			currentPlayer = this.getNextPlayer(currentPlayer, player1, player2)
		}
		err = this.makeMoveForPlayer(currentPlayer)
		if err != nil && err.Error() == NotEnoughCoinsError {
			fmt.Println(NotEnoughCoinsErrorMessage)
		} else if err != nil {
			fmt.Println(InvalidCommandErrorMessage)
		}
	}
	return this.Result(player1, player2)
}

func (this *GameRunner) makeMoveForPlayer(player1 *player.Player) error {
	commandChooseByPlayer, optionalCoinInfo, err := this.getInputFromUser(player1)
	if err != nil {
		return err
	}
	return this.Move(player1, inputCommandMap[commandChooseByPlayer], optionalCoinInfo)
}

func (this *GameRunner) getInputFromUser(player *player.Player) (string, string, error) {
	printDisplayInputMessage(player)
	return getInputCommandAndOptionCoinColor()
}

func getInputCommandAndOptionCoinColor() (string, string, error) {
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

func printDisplayInputMessage(player *player.Player) {
	fmt.Println(fmt.Sprintf("%s : Choose an outcome from the list below", player.Name()))
	fmt.Println("1. Strike")
	fmt.Println("2. Multistrike")
	fmt.Println("3. Red strike")
	fmt.Println("4. Striker strike")
	fmt.Println("5. Defunct coin")
	fmt.Println("6. None")
}

func (this *GameRunner) getNextPlayer(current *player.Player, p1 *player.Player, p2 *player.Player) *player.Player {
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
