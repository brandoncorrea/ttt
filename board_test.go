package ttt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func FullDrawBoard() [3][3]int {
	return [3][3]int{
		{-1, -1, 1},
		{1, 1, -1},
		{-1, -1, 1},
	}
}

func TestEmptyBoardIsNotGameOver(t *testing.T) {
	assert.False(t, IsGameOver(EmptyBoard()))
}

func TestFullDrawBoardIsGameOver(t *testing.T) {
	assert.True(t, IsGameOver(FullDrawBoard()))
}

func TestIncompleteGameIsNotGameOver(t *testing.T) {
	var board = FullDrawBoard()
	for row := 0; row < 3; row++ {
		for column := 0; column < 3; column++ {
			board[row][column] = 0
			assert.False(t, IsGameOver(board))
		}
	}
}

func TestWinningPlayerIsGameOver(t *testing.T) {
	for _, player := range []int{-1, 1} {
		for position := 0; position < 3; position++ {
			assert.True(t, IsGameOver(FillRow(EmptyBoard(), position, player)))
			assert.True(t, IsGameOver(FillColumn(EmptyBoard(), position, player)))
		}
		assert.True(t, IsGameOver(FillAscendingDiagonal(EmptyBoard(), player)))
		assert.True(t, IsGameOver(FillDescendingDiagonal(EmptyBoard(), player)))
	}
}

func TestNoAvailableMoves(t *testing.T) {
	assert.Empty(t, AvailableMoves(FullDrawBoard()))
}

func TestOneAvailableMove(t *testing.T) {
	for row := 0; row < 3; row++ {
		for column := 0; column < 3; column++ {
			var board = FullDrawBoard()
			board[row][column] = 0
			var expected = [2]int{row, column}
			assert.Equal(t, [][2]int{expected}, AvailableMoves(board))
		}
	}
}

func TestManyAvailableMove(t *testing.T) {
	var board = FullDrawBoard()
	var expected [][2]int
	for row := 0; row < 3; row++ {
		for column := 0; column < 3; column++ {
			board[row][column] = 0
			expected = append(expected, [2]int{row, column})
			assert.Equal(t, expected, AvailableMoves(board))
		}
	}
}

func TestNoChildrenInFullBoard(t *testing.T) {
	var board = FullDrawBoard()
	assert.Empty(t, AvailableMoves(board))
}

func TestOneChildInBoard(t *testing.T) {
	for player := range []int{-1, 1} {
		for row := 0; row < 3; row++ {
			for column := 0; column < 3; column++ {
				var board = FullDrawBoard()
				board[row][column] = 0
				var children = Children(board, player)
				board[row][column] = player
				assert.Equal(t, [][3][3]int{board}, children)
			}
		}
	}
}

func TestManyChildrenInBoard(t *testing.T) {
	var expected [][3][3]int
	for row := 0; row < 3; row++ {
		for column := 0; column < 3; column++ {
			var board = EmptyBoard()
			board[row][column] = 1
			expected = append(expected, board)
		}
	}
	assert.Equal(t, expected, Children(EmptyBoard(), 1))
}
