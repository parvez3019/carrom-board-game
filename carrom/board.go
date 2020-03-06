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

func (b *Board) HasBlackCoins(numberOfCoins int) bool {
	return b.blackCoinsOnBoard >= numberOfCoins
}

func (b *Board) HasRedCoins(numberOfCoins int) bool {
	return b.redCoinsOnBoard >= numberOfCoins
}

func (b *Board) RemoveNBlackCoins(numberOfBlackCoinsToPocket int) *Board {
	b.blackCoinsOnBoard -= numberOfBlackCoinsToPocket
	return b
}

func (b *Board) RemoveNRedCoins(numberOfRedCoinsToPocket int) *Board {
	b.redCoinsOnBoard -= numberOfRedCoinsToPocket
	return b
}

func (b *Board) HasAllCoinsExhausted() bool {
	return b.blackCoinsOnBoard == 0 && b.redCoinsOnBoard == 0
}

const (
	RED   = "red"
	BLACK = "black"
)
