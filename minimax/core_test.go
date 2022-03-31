package minimax_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/boards"
	"ttt/core"
	"ttt/it"
	"ttt/minimax"
)

var ValidBoardMoves = map[[3][3]int][][2]int{
	boards.Empty(): {{0, 0}, {0, 2}, {2, 0}, {2, 2}},
	{
		{core.AI, core.User, core.User},
		{core.User, core.AI, core.AI},
		{core.User, core.Empty, core.Empty},
	}: {{2, 2}},
	{
		{core.User, core.Empty, core.Empty},
		{core.AI, core.User, core.User},
		{core.User, core.AI, core.AI},
	}: {{0, 2}},
	{
		{core.User, core.Empty, core.Empty},
		{core.User, core.Empty, core.Empty},
		{core.Empty, core.AI, core.AI},
	}: {{2, 0}},
	{
		{core.User, core.Empty, core.Empty},
		{core.Empty, core.AI, core.User},
		{core.Empty, core.User, core.AI},
	}: {{0, 2}, {2, 0}},
	{
		{core.AI, core.User, core.Empty},
		{core.User, core.AI, core.Empty},
		{core.Empty, core.Empty, core.User},
	}: {{0, 2}, {2, 0}},
	{
		{core.Empty, core.User, core.AI},
		{core.Empty, core.AI, core.User},
		{core.User, core.Empty, core.Empty},
	}: {{0, 0}, {2, 2}},
	{
		{core.Empty, core.Empty, core.User},
		{core.User, core.AI, core.Empty},
		{core.AI, core.User, core.Empty},
	}: {{0, 0}, {2, 2}},
	{
		{core.User, core.AI, core.User},
		{core.Empty, core.AI, core.Empty},
		{core.Empty, core.User, core.Empty},
	}: {{1, 0}, {1, 2}},
	{
		{core.User, core.Empty, core.Empty},
		{core.AI, core.AI, core.User},
		{core.User, core.Empty, core.Empty},
	}: {{0, 1}, {2, 1}},
	{
		{core.Empty, core.User, core.Empty},
		{core.Empty, core.AI, core.Empty},
		{core.User, core.AI, core.User},
	}: {{1, 0}, {1, 2}},
	{
		{core.Empty, core.Empty, core.User},
		{core.User, core.AI, core.AI},
		{core.Empty, core.Empty, core.User},
	}: {{0, 1}, {2, 1}},
}

func TestOptimalMove(t *testing.T) {
	for board, validMoves := range ValidBoardMoves {
		assert.Contains(t, validMoves, minimax.OptimalMove(board))
	}
}

func TestOptimalMoveOnCompletedBoard(t *testing.T) {
	assert.Equal(t, core.BadMoveResult(), minimax.OptimalMove(boards.FullDrawBoard()))
}

func TestOptimalMoveWithOneAvailableMove(t *testing.T) {
	boards.ForIndices(func(row int, column int) {
		var board = boards.FullDrawBoard()
		board[row][column] = core.Empty
		assert.Equal(t, [2]int{row, column}, minimax.OptimalMove(board))
	})
}

func TestOptimalMoveDrawsAgainstItself(t *testing.T) {
	var board = boards.Empty()
	for turn := 0; turn < 9; turn++ {
		board = boards.AssignCell(board, minimax.OptimalMove(board), core.AI)
		boards.ForIndices(func(row int, column int) {
			board[row][column] *= -1
		})
	}

	assert.True(t, it.IsFull(board))
	assert.Equal(t, core.Empty, boards.WinningPlayer(board))
}
