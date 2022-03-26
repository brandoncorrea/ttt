package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"ttt/core"
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
	var badResult = [2]int{-1, -1}
	for _, input := range disallowedInputs {
		assert.Equal(t, badResult, ParseUserInput(input))
	}
}

func TestParsesFirstCell(t *testing.T) {
	var acceptableFormats = []string{
		"%d, %d",
		"%d %d",
		"%d,%d",
		"  \t  %d  \n  %d \r ",
	}

	core.ForIndices(func(row int, column int) {
		var cell = [2]int{row, column}
		for _, format := range acceptableFormats {
			assert.Equal(t, cell, ParseUserInput(fmt.Sprintf(format, row, column)))
		}
	})
}
