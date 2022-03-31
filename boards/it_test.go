package boards_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/boards"
	"ttt/players"
)

func TestEmptyBoardIsNotGameOver(t *testing.T) {
	assert.False(t, boards.IsGameOver(boards.Empty()))
}

func TestFullDrawBoardIsGameOver(t *testing.T) {
	assert.True(t, boards.IsGameOver(boards.FullDrawBoard()))
}

func TestIncompleteGameIsNotGameOver(t *testing.T) {
	var board = boards.FullDrawBoard()
	boards.ForIndices(func(row int, column int) {
		board[row][column] = players.Empty
		assert.False(t, boards.IsGameOver(board))
	})
}

func TestWinningPlayerIsGameOver(t *testing.T) {
	players.ForEach(func(player int) {
		for position := 0; position < 3; position++ {
			assert.True(t, boards.IsGameOver(boards.FillRow(boards.Empty(), position, player)))
			assert.True(t, boards.IsGameOver(boards.FillColumn(boards.Empty(), position, player)))
		}
		assert.True(t, boards.IsGameOver(boards.FillAscendingDiagonal(boards.Empty(), player)))
		assert.True(t, boards.IsGameOver(boards.FillDescendingDiagonal(boards.Empty(), player)))
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
	var board = boards.Empty()
	for _, move := range moves {
		assert.False(t, boards.IsValidMove(board, move))
	}
}

func TestAllInputsWithinRangeAreValid(t *testing.T) {
	var board = boards.Empty()
	boards.ForIndices(func(row int, column int) {
		assert.True(t, boards.IsValidMove(board, [2]int{row, column}))
	})
}

func TestCannotMoveToOccupiedCells(t *testing.T) {
	var board = boards.Empty()
	board[0][0] = players.AI
	board[1][1] = players.User
	assert.False(t, boards.IsValidMove(board, [2]int{}))
	assert.False(t, boards.IsValidMove(board, [2]int{1, 1}))
	assert.True(t, boards.IsValidMove(board, [2]int{0, 1}))
}
