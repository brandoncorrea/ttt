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

func ToString(board [3][3]int) string {
	var result = ""
	for row := 0; row < 3; row++ {
		for column := 0; column < 3; column++ {
			result += "| " + PlayerToString(board[row][column]) + " "
		}
		result += "|\r\n"
	}
	return strings.TrimSpace(result)
}

func AssignCell(board [3][3]int, cell [2]int, player int) [3][3]int {
	board[cell[0]][cell[1]] = player
	return board
}

func WinningPlayer(board [3][3]int) int {
	for position := 0; position < 3; position++ {
		if IsWinningRow(board, position) {
			return board[position][0]
		} else if IsWinningColumn(board, position) {
			return board[0][position]
		}
	}
	if IsWinningDiagonal(board) {
		return board[1][1]
	}
	return core.Empty
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
