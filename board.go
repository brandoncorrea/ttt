package minimax

func EmptyBoard() [3][3]int {
	return [3][3]int{
		{0, 0, 0},
		{0, 0, 0},
		{0, 0, 0}}
}

func Clone(board [3][3]int) [3][3]int {
	var newBoard [3][3]int
	for row := 0; row < 3; row++ {
		for column := 0; column < 3; column++ {
			newBoard[row][column] = board[row][column]
		}
	}
	return newBoard
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
	return section[0] != 0 &&
		section[0] == section[1] &&
		section[1] == section[2]
}

func IsWinningDiagonal(board [3][3]int) bool {
	return board[1][1] != 0 &&
		((board[1][1] == board[0][0] && board[1][1] == board[2][2]) ||
			(board[1][1] == board[0][2] && board[1][1] == board[2][0]))
}

func HasEmptyCells(board [3][3]int) bool {
	for row := 0; row < 3; row++ {
		for column := 0; column < 3; column++ {
			if board[row][column] == 0 {
				return true
			}
		}
	}
	return false
}

func IsGameOver(board [3][3]int) bool {
	for position := 0; position < 3; position++ {
		if IsWinning(Row(board, position)) ||
			IsWinning(Column(board, position)) {
			return true
		}
	}

	return IsWinningDiagonal(board) || !HasEmptyCells(board)
}

func AvailableMoves(board [3][3]int) [][2]int {
	var moves [][2]int
	for row := 0; row < 3; row++ {
		for column := 0; column < 3; column++ {
			if board[row][column] == 0 {
				moves = append(moves, [2]int{row, column})
			}
		}
	}
	return moves
}

func Children(board [3][3]int, player int) [][3][3]int {
	var children [][3][3]int
	for row := 0; row < 3; row++ {
		for column := 0; column < 3; column++ {
			if board[row][column] == 0 {
				var child = Clone(board)
				child[row][column] = player
				children = append(children, child)
			}
		}
	}
	return children
}
