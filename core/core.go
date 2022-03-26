package core

func ForIndices(f func(int, int)) {
	for row := 0; row < 3; row++ {
		for column := 0; column < 3; column++ {
			f(row, column)
		}
	}
}

func Any(board [3][3]int, pred func(int) bool) bool {
	for row := 0; row < 3; row++ {
		for column := 0; column < 3; column++ {
			if pred(board[row][column]) {
				return true
			}
		}
	}
	return false
}

func Map(board [3][3]int, f func(int) int) [3][3]int {
	ForIndices(func(row int, column int) {
		board[row][column] = f(board[row][column])
	})
	return board
}
