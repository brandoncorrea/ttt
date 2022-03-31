package game

import (
	"ttt/boards"
	"ttt/minimax"
	"ttt/players"
)

func Play(getUserMove func([3][3]int) [2]int, board [3][3]int) [3][3]int {
	for !boards.IsGameOver(board) {
		board = boards.AssignCell(board, getUserMove(board), players.User)
		if !boards.IsGameOver(board) {
			board = boards.AssignCell(board, minimax.OptimalMove(board), players.AI)
		}
	}
	return board
}
