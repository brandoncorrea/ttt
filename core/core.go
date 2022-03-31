package core

const (
	AI    = 1
	Empty = 0
	User  = -1
)

func BadMoveResult() [2]int {
	return [2]int{-1, -1}
}
