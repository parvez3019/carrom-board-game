package carrom

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func Test_ShouldReturnTrueWhenBoardHasNCoinsOfXColor(t *testing.T) {
	board := NewCarromBoard(1,1)

	assert.True(t, board.HasBlackCoins(1))
	assert.True(t, board.HasRedCoins(1))
	assert.False(t, board.HasBlackCoins(2))
	assert.False(t, board.HasRedCoins(2))
}

func Test_ShouldReturnTrueWhenBoardHasNoCoins(t *testing.T) {
	board := NewCarromBoard(0,0)

	assert.True(t, board.HasAllCoinsExhausted())
}

func Test_ShouldReturnFalseWhenBoardHasSomeCoins(t *testing.T) {
	board := NewCarromBoard(1,0)

	assert.False(t, board.HasAllCoinsExhausted())
}

func Test_ShouldReturn2BlackCoinsWhenRemoved2BlackCoinsFromBoard(t *testing.T) {
	board := NewCarromBoard(4,4)
	board.RemoveNBlackCoins(2)

	assert.True(t, board.HasBlackCoins(2))
	assert.False(t, board.HasBlackCoins(4))
	assert.True(t, board.HasRedCoins(4))
}

func Test_ShouldReturn2RedCoinsWhenRemoved2RedCoinsFromBoard(t *testing.T) {
	board := NewCarromBoard(4,4)
	board.RemoveNRedCoins(2)

	assert.True(t, board.HasRedCoins(2))
	assert.False(t, board.HasRedCoins(4))
	assert.True(t, board.HasBlackCoins(4))
}
