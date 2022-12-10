package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

var example = `
R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`

func TestSolution2(t *testing.T) {
	f, _ := os.Open("input")
	moves := ParseInput(f)
	fmt.Println(part2(moves))
}

func TestSolution1(t *testing.T) {
	f, _ := os.Open("input")
	moves := ParseInput(f)
	assert.Equal(t, 6190, part1(moves))
}

func TestParseInput(t *testing.T) {
	expected := []Move{
		{Direction: 'U', Times: 1},
		{Direction: 'L', Times: 3},
	}
	input := "U 1\nL 3\n"
	moves := ParseInput(strings.NewReader(input))
	assert.Equal(t, expected, moves)
}

func TestParsesMoves(t *testing.T) {
	up := parseMove("U 4")
	assert.Equal(t, 'U', up.Direction)
	assert.Equal(t, 4, up.Times)
	left := parseMove("L 1")
	assert.Equal(t, 'L', left.Direction)
	assert.Equal(t, 1, left.Times)
	down := parseMove("D 3")
	assert.Equal(t, 'D', down.Direction)
	assert.Equal(t, 3, down.Times)
	right := parseMove("R 2")
	assert.Equal(t, 'R', right.Direction)
	assert.Equal(t, 2, right.Times)
}
