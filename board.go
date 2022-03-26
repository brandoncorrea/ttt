package ttt

import "strings"

func EmptyBoard() [3][3]int {
	return [3][3]int{}
}

func PlayerToString(token int) string {
	if token == 1 {
		return "X"
	} else if token == -1 {
		return "O"
	} else {
		return "_"
	}
}

func BoardToString(board [3][3]int) string {
	var result = ""
	for row := 0; row < 3; row++ {
		for column := 0; column < 3; column++ {
			result += "| " + PlayerToString(board[row][column]) + " "
		}
		result += "|\r\n"
	}
	return strings.TrimSpace(result)
}

func IsEmpty(value int) bool {
	return value == 0
}

func AssignCell(board [3][3]int, cell [2]int, player int) [3][3]int {
	board[cell[0]][cell[1]] = player
	return board
}

func Row(board [3][3]int, row int) [3]int {
	return board[row]
}

func Column(board [3][3]int, column int) [3]int {
	return [3]int{
		board[0][column],
		board[1][column],
		board[2][column]}
}

func FillRow(board [3][3]int, row int, value int) [3][3]int {
	for column := 0; column < 3; column++ {
		board[row][column] = value
	}
	return board
}

func FillColumn(board [3][3]int, column int, value int) [3][3]int {
	for row := 0; row < 3; row++ {
		board[row][column] = value
	}
	return board
}

func FillDescendingDiagonal(board [3][3]int, value int) [3][3]int {
	board[0][0] = value
	board[1][1] = value
	board[2][2] = value
	return board
}

func FillAscendingDiagonal(board [3][3]int, value int) [3][3]int {
	board[0][2] = value
	board[1][1] = value
	board[2][0] = value
	return board
}

func IsWinning(section [3]int) bool {
	return !IsEmpty(section[0]) &&
		section[0] == section[1] &&
		section[1] == section[2]
}

func IsWinningDiagonal(board [3][3]int) bool {
	return !IsEmpty(board[1][1]) &&
		((board[1][1] == board[0][0] && board[1][1] == board[2][2]) ||
			(board[1][1] == board[0][2] && board[1][1] == board[2][0]))
}

func IsGameOver(board [3][3]int) bool {
	for position := 0; position < 3; position++ {
		if IsWinning(Row(board, position)) ||
			IsWinning(Column(board, position)) {
			return true
		}
	}

	return IsWinningDiagonal(board) || !Any(board, IsEmpty)
}

func AvailableMoves(board [3][3]int) [][2]int {
	var moves [][2]int
	ForIndices(func(row int, column int) {
		if IsEmpty(board[row][column]) {
			moves = append(moves, [2]int{row, column})
		}
	})
	return moves
}

func Children(board [3][3]int, player int) [][3][3]int {
	var children [][3][3]int
	for _, move := range AvailableMoves(board) {
		children = append(children, AssignCell(board, move, player))
	}
	return children
}
