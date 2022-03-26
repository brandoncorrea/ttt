package main

import (
	"fmt"
	"strings"
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

func main() {
	fmt.Println("Tic Tac Toe")
}
