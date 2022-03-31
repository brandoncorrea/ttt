package main

import (
	"fmt"
	"ttt/boards"
	"ttt/game"
	"ttt/io"
)

func main() {
	fmt.Println("Tic Tac Toe")
	fmt.Println("--------------------")
	fmt.Println(boards.ToString(game.Play(io.NewMoveReader(), boards.Empty())))
	fmt.Println("Game Over!")
}
