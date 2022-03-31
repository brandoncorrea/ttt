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

func ForEach(f func(int)) {
	for _, player := range []int{AI, User} {
		f(player)
	}
}
