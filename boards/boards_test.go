package boards

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/core"
)

func ForPlayer(f func(int)) {
	for _, player := range []int{core.AI, core.User} {
		f(player)
	}
}

func ForPlayerAndIndices(f func(int, int, int)) {
	ForPlayer(func(p int) {
		core.ForIndices(func(r int, c int) {
			f(p, r, c)
		})
	})
}

func TestEmptyBoardIsNotGameOver(t *testing.T) {
	assert.False(t, IsGameOver(EmptyBoard()))
}

func TestFullDrawBoardIsGameOver(t *testing.T) {
	assert.True(t, IsGameOver(FullDrawBoard()))
}

func TestIncompleteGameIsNotGameOver(t *testing.T) {
	var board = FullDrawBoard()
	core.ForIndices(func(row int, column int) {
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

func TestNoAvailableMoves(t *testing.T) {
	assert.Empty(t, AvailableMoves(FullDrawBoard()))
}

func TestOneAvailableMove(t *testing.T) {
	core.ForIndices(func(row int, column int) {
		var board = FullDrawBoard()
		board[row][column] = core.Empty
		var expected = [2]int{row, column}
		assert.Equal(t, [][2]int{expected}, AvailableMoves(board))
	})
}

func TestManyAvailableMove(t *testing.T) {
	var board = FullDrawBoard()
	var expected [][2]int
	core.ForIndices(func(row int, column int) {
		board[row][column] = core.Empty
		expected = append(expected, [2]int{row, column})
		assert.Equal(t, expected, AvailableMoves(board))
	})
}

func TestNoChildrenInFullBoard(t *testing.T) {
	var board = FullDrawBoard()
	assert.Empty(t, AvailableMoves(board))
}

func TestOneChildInBoard(t *testing.T) {
	ForPlayerAndIndices(func(player int, row int, column int) {
		var board = FullDrawBoard()
		board[row][column] = core.Empty
		var children = Children(board, player)
		board[row][column] = player
		assert.Equal(t, [][3][3]int{board}, children)
	})
}

func TestManyChildrenInBoard(t *testing.T) {
	ForPlayer(func(player int) {
		var expected [][3][3]int
		core.ForIndices(func(row int, column int) {
			var board = EmptyBoard()
			board[row][column] = player
			expected = append(expected, board)
		})
		assert.Equal(t, expected, Children(EmptyBoard(), player))
	})
}

func TestEmptyBoardToString(t *testing.T) {
	var expected = "| _ | _ | _ |\r\n| _ | _ | _ |\r\n| _ | _ | _ |"
	assert.Equal(t, expected, BoardToString(EmptyBoard()))
}

func TestBoardToStringWithOneMove(t *testing.T) {
	var expected = "| _ | _ | _ |\r\n| _ | _ | _ |\r\n| X | _ | _ |"
	var board = EmptyBoard()
	board[2][0] = core.AI
	assert.Equal(t, expected, BoardToString(board))
}

func TestBoardWithOneUserMove(t *testing.T) {
	var expected = "| _ | _ | _ |\r\n| _ | O | _ |\r\n| _ | _ | _ |"
	var board = EmptyBoard()
	board[1][1] = core.User
	assert.Equal(t, expected, BoardToString(board))
}

func TestBoardWithMultipleMoves(t *testing.T) {
	var expected = "| X | O | _ |\r\n| _ | X | O |\r\n| X | _ | O |"
	var board = [3][3]int{
		{1, -1, 0},
		{0, 1, -1},
		{1, 0, -1},
	}
	assert.Equal(t, expected, BoardToString(board))
}

func TestPlayerToString(t *testing.T) {
	assert.Equal(t, "X", PlayerToString(1))
	assert.Equal(t, "O", PlayerToString(-1))
	assert.Equal(t, "_", PlayerToString(0))
	assert.Equal(t, "_", PlayerToString(-2))
	assert.Equal(t, "_", PlayerToString(2))
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
	core.ForIndices(func(row int, column int) {
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
