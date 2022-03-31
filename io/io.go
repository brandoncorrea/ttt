package io

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"ttt/boards"
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
		if boards.IsValidMove(board, move) {
			return move
		}
	}
}

func NewMoveReader() func([3][3]int) [2]int {
	reader := bufio.NewReader(os.Stdin)
	return func(board [3][3]int) [2]int {
		return ReadUserMove(reader, board)
	}
}
