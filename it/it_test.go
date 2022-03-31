package it_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/boards"
	"ttt/core"
	"ttt/it"
)

func TestEmptyBoardIsNotGameOver(t *testing.T) {
	assert.False(t, it.IsGameOver(boards.Empty()))
}

func TestFullDrawBoardIsGameOver(t *testing.T) {
	assert.True(t, it.IsGameOver(boards.FullDrawBoard()))
}

func TestIncompleteGameIsNotGameOver(t *testing.T) {
	var board = boards.FullDrawBoard()
	boards.ForIndices(func(row int, column int) {
		board[row][column] = core.Empty
		assert.False(t, it.IsGameOver(board))
	})
}

func TestWinningPlayerIsGameOver(t *testing.T) {
	for _, player := range [2]int{core.User, core.AI} {
		for position := 0; position < 3; position++ {
			assert.True(t, it.IsGameOver(boards.FillRow(boards.Empty(), position, player)))
			assert.True(t, it.IsGameOver(boards.FillColumn(boards.Empty(), position, player)))
		}
		assert.True(t, it.IsGameOver(boards.FillAscendingDiagonal(boards.Empty(), player)))
		assert.True(t, it.IsGameOver(boards.FillDescendingDiagonal(boards.Empty(), player)))
	}
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
		assert.False(t, it.IsValidMove(board, move))
	}
}

func TestAllInputsWithinRangeAreValid(t *testing.T) {
	var board = boards.Empty()
	boards.ForIndices(func(row int, column int) {
		assert.True(t, it.IsValidMove(board, [2]int{row, column}))
	})
}

func TestCannotMoveToOccupiedCells(t *testing.T) {
	var board = boards.Empty()
	board[0][0] = core.AI
	board[1][1] = core.User
	assert.False(t, it.IsValidMove(board, [2]int{}))
	assert.False(t, it.IsValidMove(board, [2]int{1, 1}))
	assert.True(t, it.IsValidMove(board, [2]int{0, 1}))
}
