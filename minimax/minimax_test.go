package minimax_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/boards"
	"ttt/minimax"
	"ttt/players"
)

var ValidBoardMoves = map[[3][3]int][][2]int{
	boards.Empty(): {{0, 0}, {0, 2}, {2, 0}, {2, 2}},
	{
		{players.AI, players.User, players.User},
		{players.User, players.AI, players.AI},
		{players.User, players.Empty, players.Empty},
	}: {{2, 2}},
	{
		{players.User, players.Empty, players.Empty},
		{players.AI, players.User, players.User},
		{players.User, players.AI, players.AI},
	}: {{0, 2}},
	{
		{players.User, players.Empty, players.Empty},
		{players.User, players.Empty, players.Empty},
		{players.Empty, players.AI, players.AI},
	}: {{2, 0}},
	{
		{players.User, players.Empty, players.Empty},
		{players.Empty, players.AI, players.User},
		{players.Empty, players.User, players.AI},
	}: {{0, 2}, {2, 0}},
	{
		{players.AI, players.User, players.Empty},
		{players.User, players.AI, players.Empty},
		{players.Empty, players.Empty, players.User},
	}: {{0, 2}, {2, 0}},
	{
		{players.Empty, players.User, players.AI},
		{players.Empty, players.AI, players.User},
		{players.User, players.Empty, players.Empty},
	}: {{0, 0}, {2, 2}},
	{
		{players.Empty, players.Empty, players.User},
		{players.User, players.AI, players.Empty},
		{players.AI, players.User, players.Empty},
	}: {{0, 0}, {2, 2}},
	{
		{players.User, players.AI, players.User},
		{players.Empty, players.AI, players.Empty},
		{players.Empty, players.User, players.Empty},
	}: {{1, 0}, {1, 2}},
	{
		{players.User, players.Empty, players.Empty},
		{players.AI, players.AI, players.User},
		{players.User, players.Empty, players.Empty},
	}: {{0, 1}, {2, 1}},
	{
		{players.Empty, players.User, players.Empty},
		{players.Empty, players.AI, players.Empty},
		{players.User, players.AI, players.User},
	}: {{1, 0}, {1, 2}},
	{
		{players.Empty, players.Empty, players.User},
		{players.User, players.AI, players.AI},
		{players.Empty, players.Empty, players.User},
	}: {{0, 1}, {2, 1}},
}

func TestOptimalMove(t *testing.T) {
	for board, validMoves := range ValidBoardMoves {
		assert.Contains(t, validMoves, minimax.OptimalMove(board))
	}
}

func TestOptimalMoveOnCompletedBoard(t *testing.T) {
	assert.Equal(t, boards.BadMoveResult(), minimax.OptimalMove(boards.FullDrawBoard()))
}

func TestOptimalMoveWithOneAvailableMove(t *testing.T) {
	boards.ForIndices(func(row int, column int) {
		var board = boards.FullDrawBoard()
		board[row][column] = players.Empty
		assert.Equal(t, [2]int{row, column}, minimax.OptimalMove(board))
	})
}

func TestOptimalMoveDrawsAgainstItself(t *testing.T) {
	var board = boards.Empty()
	for turn := 0; turn < 9; turn++ {
		board = boards.AssignCell(board, minimax.OptimalMove(board), players.AI)
		boards.ForIndices(func(row int, column int) {
			board[row][column] *= -1
		})
	}

	assert.True(t, boards.IsFull(board))
	assert.Equal(t, players.Empty, boards.WinningPlayer(board))
}
