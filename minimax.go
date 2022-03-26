package minimax

import "math"

func Evaluate(board [3][3]int) float64 {
	for position := 0; position < 3; position++ {
		if IsWinning(Row(board, position)) {
			return float64(board[position][0])
		} else if IsWinning(Column(board, position)) {
			return float64(board[0][position])
		}
	}
	if IsWinningDiagonal(board) {
		return float64(board[1][1])
	}
	return 0
}

func Maximize(board [3][3]int, alpha float64, beta float64) float64 {
	var maxValue = -math.MaxFloat64
	for _, child := range Children(board, 1) {
		var value = Minimax(child, alpha, beta, false)
		maxValue = math.Max(maxValue, value)
		alpha = math.Max(alpha, value)
		if beta <= alpha {
			break
		}
	}
	return maxValue
}

func Minimize(board [3][3]int, alpha float64, beta float64) float64 {
	var minValue = math.MaxFloat64
	for _, child := range Children(board, -1) {
		var value = Minimax(child, alpha, beta, true)
		minValue = math.Min(minValue, value)
		beta = math.Min(beta, value)
		if beta <= alpha {
			break
		}
	}
	return minValue
}

func Minimax(board [3][3]int, alpha float64, beta float64, isMaximizing bool) float64 {
	if IsGameOver(board) {
		return Evaluate(board)
	} else if isMaximizing {
		return Maximize(board, alpha, beta)
	} else {
		return Minimize(board, alpha, beta)
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
	anomaly := Anomalies[board]
	if anomaly != [2]int{} {
		return anomaly
	}

	var maxMove = [2]int{-1, -1}
	var maxValue = -math.MaxFloat64
	for _, move := range AvailableMoves(board) {
		var child = Clone(board)
		child[move[0]][move[1]] = 1
		var value = Minimax(child, maxValue, math.MaxFloat64, false)
		if value >= maxValue {
			maxMove = move
			maxValue = value
		}
	}

	return maxMove
}
