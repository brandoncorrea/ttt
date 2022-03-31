package core

const (
	AI    = 1
	Empty = 0
	User  = -1
)

func BadMoveResult() [2]int {
	return [2]int{-1, -1}
}

func ForIndices(f func(int, int)) {
	for row := 0; row < 3; row++ {
		for column := 0; column < 3; column++ {
			f(row, column)
		}
	}
}
