package boards

import "ttt/players"

func BadMoveResult() [2]int {
	return [2]int{-1, -1}
}

func Empty() [3][3]int {
	return [3][3]int{}
}

func FullDrawBoard() [3][3]int {
	return [3][3]int{
		{players.User, players.User, players.AI},
		{players.AI, players.AI, players.User},
		{players.User, players.User, players.AI},
	}
}
