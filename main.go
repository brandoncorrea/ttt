package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"ttt/boards"
	"ttt/minimax"
)

func ParseUserInput(input string) [2]int {
	var fields = strings.Fields(strings.ReplaceAll(input, ",", " "))
	if len(fields) != 2 {
		return [2]int{-1, -1}
	}

	var move [2]int
	for i := range move {
		var _, err = fmt.Sscan(fields[i], &move[i])
		if err != nil {
			return [2]int{-1, -1}
		}
	}

	return move
}

func ReadUserMove(reader *bufio.Reader, board [3][3]int) [2]int {
	fmt.Println(boards.BoardToString(board))
	for {
		fmt.Print("Your Move: ")
		input, _ := reader.ReadString('\n')
		var move = ParseUserInput(input)
		if boards.IsValidMove(board, move) {
			return move
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Tic Tac Toe")
	fmt.Println("--------------------")

	var board = boards.EmptyBoard()
	for !boards.IsGameOver(board) {
		board = boards.AssignCell(board, ReadUserMove(reader, board), -1)
		if !boards.IsGameOver(board) {
			board = boards.AssignCell(board, minimax.OptimalMove(board), 1)
		}
	}

	fmt.Println(boards.BoardToString(board))
	fmt.Println("Game Over!")
}
