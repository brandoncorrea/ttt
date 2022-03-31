package boards_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/boards"
	"ttt/core"
)

func ForPlayer(f func(int)) {
	for _, player := range []int{core.AI, core.User} {
		f(player)
	}
}

func ForPlayerAndIndices(f func(int, int, int)) {
	ForPlayer(func(p int) {
		boards.ForIndices(func(r int, c int) {
			f(p, r, c)
		})
	})
}

func TestNoAvailableMoves(t *testing.T) {
	assert.Empty(t, boards.AvailableMoves(boards.FullDrawBoard()))
}

func TestOneAvailableMove(t *testing.T) {
	boards.ForIndices(func(row int, column int) {
		var board = boards.FullDrawBoard()
		board[row][column] = core.Empty
		var expected = [2]int{row, column}
		assert.Equal(t, [][2]int{expected}, boards.AvailableMoves(board))
	})
}

func TestManyAvailableMoves(t *testing.T) {
	var board = boards.FullDrawBoard()
	var expected [][2]int
	boards.ForIndices(func(row int, column int) {
		board[row][column] = core.Empty
		expected = append(expected, [2]int{row, column})
		assert.Equal(t, expected, boards.AvailableMoves(board))
	})
}

func TestNoChildrenInFullBoard(t *testing.T) {
	var board = boards.FullDrawBoard()
	assert.Empty(t, boards.AvailableMoves(board))
}

func TestOneChildInBoard(t *testing.T) {
	ForPlayerAndIndices(func(player int, row int, column int) {
		var board = boards.FullDrawBoard()
		board[row][column] = core.Empty
		var children = boards.Children(board, player)
		board[row][column] = player
		assert.Equal(t, [][3][3]int{board}, children)
	})
}

func TestManyChildrenInBoard(t *testing.T) {
	ForPlayer(func(player int) {
		var expected [][3][3]int
		boards.ForIndices(func(row int, column int) {
			var board = boards.Empty()
			board[row][column] = player
			expected = append(expected, board)
		})
		assert.Equal(t, expected, boards.Children(boards.Empty(), player))
	})
}

func TestEmptyBoardToString(t *testing.T) {
	var expected = "| _ | _ | _ |\r\n| _ | _ | _ |\r\n| _ | _ | _ |"
	assert.Equal(t, expected, boards.ToString(boards.Empty()))
}

func TestBoardWithOneAIMoveToString(t *testing.T) {
	var expected = "| _ | _ | _ |\r\n| _ | _ | _ |\r\n| X | _ | _ |"
	var board = boards.Empty()
	board[2][0] = core.AI
	assert.Equal(t, expected, boards.ToString(board))
}

func TestBoardWithOneUserMoveToString(t *testing.T) {
	var expected = "| _ | _ | _ |\r\n| _ | O | _ |\r\n| _ | _ | _ |"
	var board = boards.Empty()
	board[1][1] = core.User
	assert.Equal(t, expected, boards.ToString(board))
}

func TestBoardWithMultipleMovesToString(t *testing.T) {
	var expected = "| X | O | _ |\r\n| _ | X | O |\r\n| X | _ | O |"
	var board = [3][3]int{
		{core.AI, core.User, core.Empty},
		{core.Empty, core.AI, core.User},
		{core.AI, core.Empty, core.User},
	}
	assert.Equal(t, expected, boards.ToString(board))
}

func TestPlayerToString(t *testing.T) {
	assert.Equal(t, "X", boards.PlayerToString(core.AI))
	assert.Equal(t, "O", boards.PlayerToString(core.User))
	assert.Equal(t, "_", boards.PlayerToString(core.Empty))
	assert.Equal(t, "_", boards.PlayerToString(-2))
	assert.Equal(t, "_", boards.PlayerToString(2))
}

func TestNoWinningPlayerOnEmptyBoard(t *testing.T) {
	assert.Equal(t, core.Empty, boards.WinningPlayer(boards.Empty()))
}

func TestWinningPlayerOnRowColumnOrDiagonal(t *testing.T) {
	for player := range []int{core.User, core.AI} {
		var emptyBoard = boards.Empty()
		for position := 0; position < 3; position++ {
			assert.Equal(t, player, boards.WinningPlayer(boards.FillRow(emptyBoard, position, player)))
			assert.Equal(t, player, boards.WinningPlayer(boards.FillColumn(emptyBoard, position, player)))
		}

		assert.Equal(t, player, boards.WinningPlayer(boards.FillDescendingDiagonal(emptyBoard, player)))
		assert.Equal(t, player, boards.WinningPlayer(boards.FillAscendingDiagonal(emptyBoard, player)))
	}
}

func TestNoWinnerOnDrawBoard(t *testing.T) {
	assert.Equal(t, core.Empty, boards.WinningPlayer(boards.FullDrawBoard()))
}
