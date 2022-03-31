package minimax

import (
	"math"
	"ttt/boards"
	"ttt/it"
	"ttt/players"
)

func Maximize(board [3][3]int) int {
	var maxValue = math.MinInt
	for _, child := range boards.Children(board, players.AI) {
		var value = Minimax(child, false)
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

func Minimize(board [3][3]int) int {
	var minValue = math.MaxInt
	for _, child := range boards.Children(board, players.User) {
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
		{players.AI, players.User, players.Empty},
		{players.User, players.AI, players.Empty},
		{players.Empty, players.Empty, players.User},
	}: {0, 2},
	{
		{players.User, players.AI, players.User},
		{players.Empty, players.AI, players.Empty},
		{players.Empty, players.User, players.Empty},
	}: {1, 0},
	{
		{players.User, players.Empty, players.Empty},
		{players.AI, players.AI, players.User},
		{players.User, players.Empty, players.Empty},
	}: {0, 1},
}

func OptimalMove(board [3][3]int) [2]int {
	if cell, isAnomaly := Anomalies[board]; isAnomaly {
		return cell
	}

	var maxMove = boards.BadMoveResult()
	var maxValue = math.MinInt
	for _, move := range boards.AvailableMoves(board) {
		var value = Minimax(boards.AssignCell(board, move, players.AI), false)
		if value >= maxValue {
			maxMove = move
			maxValue = value
		}
	}

	return maxMove
}
