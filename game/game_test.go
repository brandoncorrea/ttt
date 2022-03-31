package game_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/boards"
	"ttt/game"
	"ttt/minimax"
	"ttt/players"
)

func FirstEmptyCell(board [3][3]int) [2]int {
	for row := 0; row < 3; row++ {
		for column := 0; column < 3; column++ {
			if boards.IsEmpty(board[row][column]) {
				return [2]int{row, column}
			}
		}
	}
	return boards.BadMoveResult()
}

func StackMoveSelector(moves [][2]int) func([3][3]int) [2]int {
	var index = -1
	return func(_ [3][3]int) [2]int {
		index++
		return moves[index]
	}
}

func TestPlayVersusMinimax(t *testing.T) {
	board := game.Play(minimax.OptimalMove, boards.Empty())
	assert.True(t, boards.IsGameOver(board))
	assert.Equal(t, players.AI, boards.WinningPlayer(board))
}

func TestChooseFirstEmptyCell(t *testing.T) {
	board := game.Play(FirstEmptyCell, boards.Empty())
	assert.True(t, boards.IsGameOver(board))
	assert.Equal(t, players.AI, boards.WinningPlayer(board))
}

func TestDrawAgainstAI(t *testing.T) {
	var getMove = StackMoveSelector([][2]int{
		{0, 0},
		{2, 0},
		{1, 2},
		{2, 1},
		{0, 2},
	})
	board := game.Play(getMove, boards.Empty())
	assert.True(t, boards.IsGameOver(board))
	assert.Equal(t, players.Empty, boards.WinningPlayer(board))
}
