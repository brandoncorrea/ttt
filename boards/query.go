package boards

import "ttt/core"

func IsEmpty(value int) bool {
	return value == core.Empty
}

func IsValidMove(board [3][3]int, move [2]int) bool {
	var row = move[0]
	var column = move[1]
	return 0 <= row && row <= 2 &&
		0 <= column && column <= 2 &&
		IsEmpty(board[row][column])
}

func IsFull(board [3][3]int) bool {
	for row := 0; row < 3; row++ {
		for column := 0; column < 3; column++ {
			if IsEmpty(board[row][column]) {
				return false
			}
		}
	}
	return true
}

func IsWinningDiagonal(board [3][3]int) bool {
	return !IsEmpty(board[1][1]) &&
		((board[1][1] == board[0][0] && board[1][1] == board[2][2]) ||
			(board[1][1] == board[0][2] && board[1][1] == board[2][0]))
}

func IsWinningRow(board [3][3]int, row int) bool {
	return !IsEmpty(board[row][0]) &&
		board[row][0] == board[row][1] &&
		board[row][1] == board[row][2]
}

func IsWinningColumn(board [3][3]int, column int) bool {
	return !IsEmpty(board[0][column]) &&
		board[0][column] == board[1][column] &&
		board[1][column] == board[2][column]
}

func IsGameOver(board [3][3]int) bool {
	for position := 0; position < 3; position++ {
		if IsWinningRow(board, position) ||
			IsWinningColumn(board, position) {
			return true
		}
	}

	return IsWinningDiagonal(board) || IsFull(board)
}
