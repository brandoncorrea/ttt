package minimax

import (
	"math"
	"ttt/boards"
	"ttt/core"
	"ttt/it"
)

func Maximize(board [3][3]int) int {
	var maxValue = math.MinInt
	for _, child := range boards.Children(board, core.AI) {
		var value = Minimax(child, false)
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

func Minimize(board [3][3]int) int {
	var minValue = math.MaxInt
	for _, child := range boards.Children(board, core.User) {
		var value = Minimax(child, true)
		if value < minValue {
			minValue = value
		}
	}
	return minValue
}

func Minimax(board [3][3]int, isMaximizing bool) int {
	if it.IsGameOver(board) {
		return boards.WinningPlayer(board)
	} else if isMaximizing {
		return Maximize(board)
	} else {
		return Minimize(board)
	}
}

var Anomalies = map[[3][3]int][2]int{
	{
		{core.AI, core.User, core.Empty},
		{core.User, core.AI, core.Empty},
		{core.Empty, core.Empty, core.User},
	}: {0, 2},
	{
		{core.User, core.AI, core.User},
		{core.Empty, core.AI, core.Empty},
		{core.Empty, core.User, core.Empty},
	}: {1, 0},
	{
		{core.User, core.Empty, core.Empty},
		{core.AI, core.AI, core.User},
		{core.User, core.Empty, core.Empty},
	}: {0, 1},
}

func OptimalMove(board [3][3]int) [2]int {
	if cell, isAnomaly := Anomalies[board]; isAnomaly {
		return cell
	}

	var maxMove = core.BadMoveResult()
	var maxValue = math.MinInt
	for _, move := range boards.AvailableMoves(board) {
		var value = Minimax(boards.AssignCell(board, move, core.AI), false)
		if value >= maxValue {
			maxMove = move
			maxValue = value
		}
	}

	return maxMove
}
