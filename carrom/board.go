package carrom

type Board struct {
	blackCoinsOnBoard int
	redCoinsOnBoard   int
}

func NewCarromBoard(numberOfBlackCoins int, numberOfRedCoins int) *Board {
	return &Board{
		blackCoinsOnBoard: numberOfBlackCoins,
		redCoinsOnBoard:   numberOfRedCoins,
	}
}

func (this *Board) HasBlackCoins(numberOfCoins int) bool {
	return this.blackCoinsOnBoard >= numberOfCoins
}

func (this *Board) HasRedCoins(numberOfCoins int) bool {
	return this.redCoinsOnBoard >= numberOfCoins
}

func (this *Board) PocketBlackCoins(numberOfBlackCoinsToPocket int) *Board {
	this.blackCoinsOnBoard -= numberOfBlackCoinsToPocket
	return this
}

func (this *Board) PocketRedCoins(numberOfRedCoinsToPocket int) *Board {
	this.redCoinsOnBoard -= numberOfRedCoinsToPocket
	return this
}

func (this *Board) HasAllCoinsExhausted() bool {
	return this.blackCoinsOnBoard == 0 && this.redCoinsOnBoard == 0
}

const (
	RED   = "red"
	BLACK = "black"
)
