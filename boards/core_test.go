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
	assert.Equal(t, expected, ToString(EmptyBoard()))
}

func TestBoardToStringWithOneMove(t *testing.T) {
	var expected = "| _ | _ | _ |\r\n| _ | _ | _ |\r\n| X | _ | _ |"
	var board = EmptyBoard()
	board[2][0] = core.AI
	assert.Equal(t, expected, ToString(board))
}

func TestBoardWithOneUserMove(t *testing.T) {
	var expected = "| _ | _ | _ |\r\n| _ | O | _ |\r\n| _ | _ | _ |"
	var board = EmptyBoard()
	board[1][1] = core.User
	assert.Equal(t, expected, ToString(board))
}

func TestBoardWithMultipleMoves(t *testing.T) {
	var expected = "| X | O | _ |\r\n| _ | X | O |\r\n| X | _ | O |"
	var board = [3][3]int{
		{core.AI, core.User, core.Empty},
		{core.Empty, core.AI, core.User},
		{core.AI, core.Empty, core.User},
	}
	assert.Equal(t, expected, ToString(board))
}

func TestPlayerToString(t *testing.T) {
	assert.Equal(t, "X", PlayerToString(core.AI))
	assert.Equal(t, "O", PlayerToString(core.User))
	assert.Equal(t, "_", PlayerToString(core.Empty))
	assert.Equal(t, "_", PlayerToString(-2))
	assert.Equal(t, "_", PlayerToString(2))
}

func TestNoWinningPlayerOnEmptyBoard(t *testing.T) {
	assert.Equal(t, core.Empty, WinningPlayer(EmptyBoard()))
}

func TestWinningPlayerOnRowColumnOrDiagonal(t *testing.T) {
	for player := range []int{core.User, core.AI} {
		var emptyBoard = EmptyBoard()
		for position := 0; position < 3; position++ {
			assert.Equal(t, player, WinningPlayer(FillRow(emptyBoard, position, player)))
			assert.Equal(t, player, WinningPlayer(FillColumn(emptyBoard, position, player)))
		}

		assert.Equal(t, player, WinningPlayer(FillDescendingDiagonal(emptyBoard, player)))
		assert.Equal(t, player, WinningPlayer(FillAscendingDiagonal(emptyBoard, player)))
	}
}

func TestNoWinnerOnDrawBoard(t *testing.T) {
	assert.Equal(t, core.Empty, WinningPlayer(FullDrawBoard()))
}
