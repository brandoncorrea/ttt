package boards

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
