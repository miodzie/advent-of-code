package main

import (
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

func TestSolutionPart1(t *testing.T) {
	f, _ := os.Open("input")
	moves := ParseInput(f)
	rope := &Rope{}
	ans := AllMoves(rope, moves)
	assert.Equal(t, 6190, ans)
}

func TestExample1(t *testing.T) {
	moves := ParseInput(strings.NewReader(example))
	rope := &Rope{}

	assert.Equal(t, 13, AllMoves(rope, moves))
}

func TestApplyMove_LEFT_3(t *testing.T) {
	rope := Rope{
		Head: Point{4, 4},
		Tail: Point{4, 3}}
	move := Move{LEFT, 3}

	tailPoints := applyMove(&rope, move)

	assert.NotEmpty(t, tailPoints)
	assert.Len(t, tailPoints, 3)
	assert.Equal(t, Point{1, 4}, rope.Head)
	assert.Equal(t, Point{2, 4}, rope.Tail)
}

func TestApplyMove(t *testing.T) {
	rope := Rope{
		Head: Point{4, 1},
		Tail: Point{3, 0}}
	move := Move{UP, 1}

	tailPoints := applyMove(&rope, move)

	assert.Equal(t, Point{4, 2}, rope.Head)
	assert.Equal(t, Point{4, 1}, rope.Tail)
	assert.Len(t, tailPoints, 2)
	assert.Equal(t, Point{3, 0}, tailPoints[0])
	assert.Equal(t, Point{4, 1}, tailPoints[1])
}

func TestParseInput(t *testing.T) {
	expected := []Move{
		{Direction: UP, Times: 1},
		{Direction: LEFT, Times: 3},
	}
	input := "U 1\nL 3\n"
	moves := ParseInput(strings.NewReader(input))
	assert.Equal(t, expected, moves)
}

func TestParsesMoves(t *testing.T) {
	up := parseMove("U 4")
	assert.Equal(t, UP, up.Direction)
	assert.Equal(t, 4, up.Times)
	left := parseMove("L 1")
	assert.Equal(t, LEFT, left.Direction)
	assert.Equal(t, 1, left.Times)
	down := parseMove("D 3")
	assert.Equal(t, DOWN, down.Direction)
	assert.Equal(t, 3, down.Times)
	right := parseMove("R 2")
	assert.Equal(t, RIGHT, right.Direction)
	assert.Equal(t, 2, right.Times)
}
