package boards_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/boards"
	"ttt/players"
)

func TestNoAvailableMoves(t *testing.T) {
	assert.Empty(t, boards.AvailableMoves(boards.FullDrawBoard()))
}

func TestOneAvailableMove(t *testing.T) {
	boards.ForIndices(func(row int, column int) {
		var board = boards.FullDrawBoard()
		board[row][column] = players.Empty
		var expected = [2]int{row, column}
		assert.Equal(t, [][2]int{expected}, boards.AvailableMoves(board))
	})
}

func TestManyAvailableMoves(t *testing.T) {
	var board = boards.FullDrawBoard()
	var expected [][2]int
	boards.ForIndices(func(row int, column int) {
		board[row][column] = players.Empty
		expected = append(expected, [2]int{row, column})
		assert.Equal(t, expected, boards.AvailableMoves(board))
	})
}

func TestNoChildrenInFullBoard(t *testing.T) {
	var board = boards.FullDrawBoard()
	assert.Empty(t, boards.AvailableMoves(board))
}

func TestOneChildInBoard(t *testing.T) {
	players.ForEach(func(player int) {
		boards.ForIndices(func(row int, column int) {
			var board = boards.FullDrawBoard()
			board[row][column] = players.Empty
			var children = boards.Children(board, player)
			board[row][column] = player
			assert.Equal(t, [][3][3]int{board}, children)
		})
	})
}

func TestManyChildrenInBoard(t *testing.T) {
	players.ForEach(func(player int) {
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
	board[2][0] = players.AI
	assert.Equal(t, expected, boards.ToString(board))
}

func TestBoardWithOneUserMoveToString(t *testing.T) {
	var expected = "| _ | _ | _ |\r\n| _ | O | _ |\r\n| _ | _ | _ |"
	var board = boards.Empty()
	board[1][1] = players.User
	assert.Equal(t, expected, boards.ToString(board))
}

func TestBoardWithMultipleMovesToString(t *testing.T) {
	var expected = "| X | O | _ |\r\n| _ | X | O |\r\n| X | _ | O |"
	var board = [3][3]int{
		{players.AI, players.User, players.Empty},
		{players.Empty, players.AI, players.User},
		{players.AI, players.Empty, players.User},
	}
	assert.Equal(t, expected, boards.ToString(board))
}

func TestNoWinningPlayerOnEmptyBoard(t *testing.T) {
	assert.Equal(t, players.Empty, boards.WinningPlayer(boards.Empty()))
}

func TestWinningPlayerOnRowColumnOrDiagonal(t *testing.T) {
	players.ForEach(func(player int) {
		var emptyBoard = boards.Empty()
		for position := 0; position < 3; position++ {
			assert.Equal(t, player, boards.WinningPlayer(boards.FillRow(emptyBoard, position, player)))
			assert.Equal(t, player, boards.WinningPlayer(boards.FillColumn(emptyBoard, position, player)))
		}

		assert.Equal(t, player, boards.WinningPlayer(boards.FillDescendingDiagonal(emptyBoard, player)))
		assert.Equal(t, player, boards.WinningPlayer(boards.FillAscendingDiagonal(emptyBoard, player)))
	})
}

func TestNoWinnerOnDrawBoard(t *testing.T) {
	assert.Equal(t, players.Empty, boards.WinningPlayer(boards.FullDrawBoard()))
}
