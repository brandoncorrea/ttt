package minimax

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/boards"
	"ttt/core"
)

func TestEvaluatesEmptyBoard(t *testing.T) {
	assert.Equal(t, 0, Evaluate(boards.EmptyBoard()))
}

func TestEvaluatesWinningRowColumnOrDiagonal(t *testing.T) {
	for player := range []int{-1, 1} {
		var emptyBoard = boards.EmptyBoard()
		for position := 0; position < 3; position++ {
			assert.Equal(t, player, Evaluate(boards.FillRow(emptyBoard, position, player)))
			assert.Equal(t, player, Evaluate(boards.FillColumn(emptyBoard, position, player)))
		}

		assert.Equal(t, player, Evaluate(boards.FillDescendingDiagonal(emptyBoard, player)))
		assert.Equal(t, player, Evaluate(boards.FillAscendingDiagonal(emptyBoard, player)))
	}
}

func TestEvaluatesDrawBoard(t *testing.T) {
	var board = [3][3]int{
		{-1, -1, 1},
		{1, -1, -1},
		{-1, 1, 1},
	}
	assert.Equal(t, 0, Evaluate(board))
}

func TestOptimalMoveOnCompletedBoard(t *testing.T) {
	assert.Equal(t, [2]int{-1, -1}, OptimalMove(boards.FullDrawBoard()))
}

func TestOptimalMoveOnEmptyBoardChoosesCorner(t *testing.T) {
	var expected = [][2]int{
		{0, 0},
		{0, 2},
		{2, 0},
		{2, 2},
	}
	assert.Contains(t, expected, OptimalMove(boards.EmptyBoard()))
}

func TestOptimalMoveWithOneAvailableMove(t *testing.T) {
	core.ForIndices(func(row int, column int) {
		var board = [3][3]int{
			{-1, -1, 1},
			{1, -1, -1},
			{-1, 1, 1},
		}
		board[row][column] = 0
		assert.Equal(t, [2]int{row, column}, OptimalMove(board))
	})
}

func TestOptimalMoveChoosesWinOverDraw(t *testing.T) {
	var board = [3][3]int{
		{1, -1, -1},
		{-1, 1, 1},
		{-1, 0, 0},
	}

	assert.Equal(t, [2]int{2, 2}, OptimalMove(board))
}

func TestOptimalMoveChoosesDrawOverLoss(t *testing.T) {
	var board = [3][3]int{
		{-1, 0, 0},
		{1, -1, -1},
		{-1, 1, 1},
	}

	assert.Equal(t, [2]int{0, 2}, OptimalMove(board))
}

func TestOptimalMoveChoosesWinOverLoss(t *testing.T) {
	var board = [3][3]int{
		{-1, 0, 0},
		{-1, 0, 0},
		{0, 1, 1},
	}

	assert.Equal(t, [2]int{2, 0}, OptimalMove(board))
}

func TestOptimalMoveForBottomRightDipper(t *testing.T) {
	var board = [3][3]int{
		{-1, 0, 0},
		{0, 1, -1},
		{0, -1, 1},
	}
	assert.Contains(t, [][2]int{{0, 2}, {2, 0}}, OptimalMove(board))
}

func TestOptimalMoveForTopLeftDipper(t *testing.T) {
	var board = [3][3]int{
		{1, -1, 0},
		{-1, 1, 0},
		{0, 0, -1},
	}
	assert.Contains(t, [][2]int{{0, 2}, {2, 0}}, OptimalMove(board))
}

func TestOptimalMoveForTopRightDipper(t *testing.T) {
	var board = [3][3]int{
		{0, -1, 1},
		{0, 1, -1},
		{-1, 0, 0},
	}
	assert.Contains(t, [][2]int{{0, 0}, {2, 2}}, OptimalMove(board))
}

func TestOptimalMoveForBottomLeftDipper(t *testing.T) {
	var board = [3][3]int{
		{0, 0, -1},
		{-1, 1, 0},
		{1, -1, 0},
	}
	assert.Contains(t, [][2]int{{0, 0}, {2, 2}}, OptimalMove(board))
}

func TestOptimalMoveForTopT(t *testing.T) {
	var board = [3][3]int{
		{-1, 1, -1},
		{0, 1, 0},
		{0, -1, 0},
	}
	assert.Contains(t, [][2]int{{1, 0}, {1, 2}}, OptimalMove(board))
}

func TestOptimalMoveForLeftT(t *testing.T) {
	var board = [3][3]int{
		{-1, 0, 0},
		{1, 1, -1},
		{-1, 0, 0},
	}
	assert.Contains(t, [][2]int{{0, 1}, {2, 1}}, OptimalMove(board))
}

func TestOptimalMoveForBottomT(t *testing.T) {
	var board = [3][3]int{
		{0, -1, 0},
		{0, 1, 0},
		{-1, 1, -1},
	}
	assert.Contains(t, [][2]int{{1, 0}, {1, 2}}, OptimalMove(board))
}

func TestOptimalMoveForRightT(t *testing.T) {
	var board = [3][3]int{
		{0, 0, -1},
		{-1, 1, 1},
		{0, 0, -1},
	}
	assert.Contains(t, [][2]int{{0, 1}, {2, 1}}, OptimalMove(board))
}

func TestOptimalMoveDrawsAgainstItself(t *testing.T) {
	var negate = func(x int) int { return x * -1 }
	var board = boards.EmptyBoard()
	for turn := 0; turn < 9; turn++ {
		board = core.Map(boards.AssignCell(board, OptimalMove(board), 1), negate)
	}

	var expected = [3][3]int{
		{-1, 1, -1},
		{-1, 1, 1},
		{1, -1, -1},
	}
	assert.Equal(t, expected, board)
}
