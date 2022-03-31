package minimax

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/boards"
	"ttt/core"
)

var ValidBoardMoves = map[[3][3]int][][2]int{
	boards.EmptyBoard(): {{0, 0}, {0, 2}, {2, 0}, {2, 2}},
	{
		{1, -1, -1},
		{-1, 1, 1},
		{-1, 0, 0},
	}: {{2, 2}},
	{
		{-1, 0, 0},
		{1, -1, -1},
		{-1, 1, 1},
	}: {{0, 2}},
	{
		{-1, 0, 0},
		{-1, 0, 0},
		{0, 1, 1},
	}: {{2, 0}},
	{
		{-1, 0, 0},
		{0, 1, -1},
		{0, -1, 1},
	}: {{0, 2}, {2, 0}},
	{
		{1, -1, 0},
		{-1, 1, 0},
		{0, 0, -1},
	}: {{0, 2}, {2, 0}},
	{
		{0, -1, 1},
		{0, 1, -1},
		{-1, 0, 0},
	}: {{0, 0}, {2, 2}},
	{
		{0, 0, -1},
		{-1, 1, 0},
		{1, -1, 0},
	}: {{0, 0}, {2, 2}},
	{
		{-1, 1, -1},
		{0, 1, 0},
		{0, -1, 0},
	}: {{1, 0}, {1, 2}},
	{
		{-1, 0, 0},
		{1, 1, -1},
		{-1, 0, 0},
	}: {{0, 1}, {2, 1}},
	{
		{0, -1, 0},
		{0, 1, 0},
		{-1, 1, -1},
	}: {{1, 0}, {1, 2}},
	{
		{0, 0, -1},
		{-1, 1, 1},
		{0, 0, -1},
	}: {{0, 1}, {2, 1}},
}

func TestOptimalMove(t *testing.T) {
	for board, validMoves := range ValidBoardMoves {
		assert.Contains(t, validMoves, OptimalMove(board))
	}
}

func TestOptimalMoveOnCompletedBoard(t *testing.T) {
	assert.Equal(t, core.BadMoveResult(), OptimalMove(boards.FullDrawBoard()))
}

func TestOptimalMoveWithOneAvailableMove(t *testing.T) {
	boards.ForIndices(func(row int, column int) {
		var board = boards.FullDrawBoard()
		board[row][column] = core.Empty
		assert.Equal(t, [2]int{row, column}, OptimalMove(board))
	})
}

func TestOptimalMoveDrawsAgainstItself(t *testing.T) {
	var board = boards.EmptyBoard()
	for turn := 0; turn < 9; turn++ {
		board = boards.AssignCell(board, OptimalMove(board), core.AI)
		boards.ForIndices(func(row int, column int) {
			board[row][column] *= -1
		})
	}

	assert.True(t, boards.IsFull(board))
	assert.Equal(t, core.Empty, boards.WinningPlayer(board))
}