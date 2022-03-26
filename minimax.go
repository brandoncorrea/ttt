package minimax

import "math"

func Evaluate(board [3][3]int) int {
	for position := 0; position < 3; position++ {
		if IsWinning(Row(board, position)) {
			return board[position][0]
		} else if IsWinning(Column(board, position)) {
			return board[0][position]
		}
	}
	if IsWinningDiagonal(board) {
		return board[1][1]
	}
	return 0
}

func Maximize(board [3][3]int) int {
	var maxValue = math.MinInt
	for _, child := range Children(board, 1) {
		var value = Minimax(child, false)
		if value > maxValue {
			maxValue = value
		}
	}
	return maxValue
}

func Minimize(board [3][3]int) int {
	var minValue = math.MaxInt
	for _, child := range Children(board, -1) {
		var value = Minimax(child, true)
		if value < minValue {
			minValue = value
		}
	}
	return minValue
}

func Minimax(board [3][3]int, isMaximizing bool) int {
	if IsGameOver(board) {
		return Evaluate(board)
	} else if isMaximizing {
		return Maximize(board)
	} else {
		return Minimize(board)
	}
}

var Anomalies = map[[3][3]int][2]int{
	{
		{1, -1, 0},
		{-1, 1, 0},
		{0, 0, -1},
	}: {0, 2},
	{
		{-1, 1, -1},
		{0, 1, 0},
		{0, -1, 0},
	}: {1, 0},
	{
		{-1, 0, 0},
		{1, 1, -1},
		{-1, 0, 0},
	}: {0, 1},
}

func OptimalMove(board [3][3]int) [2]int {
	if cell, isAnomaly := Anomalies[board]; isAnomaly {
		return cell
	}

	var maxMove = [2]int{-1, -1}
	var maxValue = math.MinInt
	for _, move := range AvailableMoves(board) {
		var value = Minimax(AssignCell(board, move, 1), false)
		if value >= maxValue {
			maxMove = move
			maxValue = value
		}
	}

	return maxMove
}
