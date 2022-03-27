package boards

import (
	"strings"
	"ttt/core"
)

func EmptyBoard() [3][3]int {
	return [3][3]int{}
}

func FullDrawBoard() [3][3]int {
	return [3][3]int{
		{core.User, core.User, core.AI},
		{core.AI, core.AI, core.User},
		{core.User, core.User, core.AI},
	}
}

func PlayerToString(token int) string {
	if token == core.AI {
		return "X"
	} else if token == core.User {
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
	return value == core.Empty
}

func IsValidMove(board [3][3]int, move [2]int) bool {
	var row = move[0]
	var column = move[1]
	return 0 <= row && row <= 2 &&
		0 <= column && column <= 2 &&
		IsEmpty(board[row][column])
}

func AssignCell(board [3][3]int, cell [2]int, player int) [3][3]int {
	board[cell[0]][cell[1]] = player
	return board
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

func IsBoardFull(board [3][3]int) bool {
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

	return IsWinningDiagonal(board) || IsBoardFull(board)
}

func AvailableMoves(board [3][3]int) [][2]int {
	var moves [][2]int
	core.ForIndices(func(row int, column int) {
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
