# This project has been moved to [tictactoe](https://github.com/brandoncorrea/tictactoe/tree/main/go)

# ttt

Tic-Tac-Toe game written in Go.

### Objective

Take a vertical, horizontal, or diagonal row to win!

### Moving

Enter two 0-indexed numbers to take a square.

#### Examples

````
Your Move: 0 1
| _ | O | _ |
| _ | _ | _ |
| _ | _ | _ |
````

````
Your Move: 1 0
| _ | _ | _ |
| O | _ | _ |
| _ | _ | _ |
````

````
Your Move: 2 2
| _ | _ | _ |
| _ | _ | _ |
| _ | _ | O |
````

### Execute
`go run main.go`

### Test

#### Run all Tests
`go test ./...`

#### Run Tests Verbose
`go test -v ./...`

#### Autorun Tests Verbose
`/bin/zsh ./autotest.sh`

