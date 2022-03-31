package players_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/players"
)

func TestPlayerToString(t *testing.T) {
	assert.Equal(t, "X", players.ToString(players.AI))
	assert.Equal(t, "O", players.ToString(players.User))
	assert.Equal(t, "_", players.ToString(players.Empty))
	assert.Equal(t, "_", players.ToString(-2))
	assert.Equal(t, "_", players.ToString(2))
}
