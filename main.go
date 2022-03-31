package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"ttt/boards"
	"ttt/it"
	"ttt/minimax"
	"ttt/players"
)

func ParseUserInput(input string) [2]int {
	var fields = strings.Fields(strings.ReplaceAll(input, ",", " "))
	if len(fields) != 2 {
		return boards.BadMoveResult()
	}

	var move [2]int
	for i := range move {
		var _, err = fmt.Sscan(fields[i], &move[i])
		if err != nil {
			return boards.BadMoveResult()
		}
	}

	return move
}

func ReadUserMove(reader *bufio.Reader, board [3][3]int) [2]int {
	fmt.Println(boards.ToString(board))
	for {
		fmt.Print("Your Move: ")
		input, _ := reader.ReadString('\n')
		var move = ParseUserInput(input)
		if it.IsValidMove(board, move) {
			return move
		}
	}
}

func Play(reader *bufio.Reader, board [3][3]int) [3][3]int {
	for !it.IsGameOver(board) {
		board = boards.AssignCell(board, ReadUserMove(reader, board), players.User)
		if !it.IsGameOver(board) {
			board = boards.AssignCell(board, minimax.OptimalMove(board), players.AI)
		}
	}
	return board
}

func main() {
	fmt.Println("Tic Tac Toe")
	fmt.Println("--------------------")
	fmt.Println(boards.ToString(Play(bufio.NewReader(os.Stdin), boards.Empty())))
	fmt.Println("Game Over!")
}
