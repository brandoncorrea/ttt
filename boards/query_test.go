package boards

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/core"
)

func TestEmptyBoardIsNotGameOver(t *testing.T) {
	assert.False(t, IsGameOver(EmptyBoard()))
}

func TestFullDrawBoardIsGameOver(t *testing.T) {
	assert.True(t, IsGameOver(FullDrawBoard()))
}

func TestIncompleteGameIsNotGameOver(t *testing.T) {
	var board = FullDrawBoard()
	ForIndices(func(row int, column int) {
		board[row][column] = core.Empty
		assert.False(t, IsGameOver(board))
	})
}

func TestWinningPlayerIsGameOver(t *testing.T) {
	ForPlayer(func(player int) {
		for position := 0; position < 3; position++ {
			assert.True(t, IsGameOver(FillRow(EmptyBoard(), position, player)))
			assert.True(t, IsGameOver(FillColumn(EmptyBoard(), position, player)))
		}
		assert.True(t, IsGameOver(FillAscendingDiagonal(EmptyBoard(), player)))
		assert.True(t, IsGameOver(FillDescendingDiagonal(EmptyBoard(), player)))
	})
}

func TestOutOfRangeMovesAreInvalid(t *testing.T) {
	var moves = [][2]int{
		{-1, -1},
		{-1, 1},
		{1, -1},
		{3, 1},
		{1, 3},
		{3, 3},
	}
	var board = EmptyBoard()
	for _, move := range moves {
		assert.False(t, IsValidMove(board, move))
	}
}

func TestAllInputsWithinRangeAreValid(t *testing.T) {
	var board = EmptyBoard()
	ForIndices(func(row int, column int) {
		assert.True(t, IsValidMove(board, [2]int{row, column}))
	})
}

func TestCannotMoveToOccupiedCells(t *testing.T) {
	var board = EmptyBoard()
	board[0][0] = core.AI
	board[1][1] = core.User
	assert.False(t, IsValidMove(board, [2]int{}))
	assert.False(t, IsValidMove(board, [2]int{1, 1}))
	assert.True(t, IsValidMove(board, [2]int{0, 1}))
}
