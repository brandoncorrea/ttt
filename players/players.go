package players

const (
	AI    = 1
	Empty = 0
	User  = -1
)

func ToString(token int) string {
	if token == AI {
		return "X"
	} else if token == User {
		return "O"
	} else {
		return "_"
	}
}
