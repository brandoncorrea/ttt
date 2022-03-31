package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/boards"
)

func TestParsesBadInput(t *testing.T) {
	var disallowedInputs = []string{
		"",
		"\n",
		"\r\n",
		" ",
		"a",
		"a b",
		"a b c",
		"1 a",
		"a 1",
		"1",
		"1 2 3",
		"1 - 3",
	}

	for _, input := range disallowedInputs {
		assert.Equal(t, boards.BadMoveResult(), ParseUserInput(input))
	}
}

func TestParsesFirstCell(t *testing.T) {
	var acceptableFormats = []string{
		"%d, %d",
		"%d %d",
		"%d,%d",
		"  \t  %d  \n  %d \r ",
	}

	boards.ForIndices(func(row int, column int) {
		var cell = [2]int{row, column}
		for _, format := range acceptableFormats {
			assert.Equal(t, cell, ParseUserInput(fmt.Sprintf(format, row, column)))
		}
	})
}
