package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestRope_ApplyMovePullsRopeBehind_D_1(t *testing.T) {
	rope := Rope{
		Head: Point{1, 3},
		Tail: Point{2, 4}}
	move := Move{Direction: DOWN, Times: 1}
	rope.Apply(move)
	assert.Equal(t, Point{1, 2}, rope.Head)
	assert.Equal(t, Point{1, 3}, rope.Tail)
}

func TestRope_ApplyMovePullsRopeBehind_L_5(t *testing.T) {
	rope := Rope{
		Head: Point{5, 2},
		Tail: Point{4, 3}}
	move := Move{Direction: LEFT, Times: 5}
	rope.Apply(move)
	assert.Equal(t, Point{0, 2}, rope.Head)
	assert.Equal(t, Point{1, 2}, rope.Tail)
}

func TestRope_ApplyMovePullsRopeBehind_UP_4(t *testing.T) {
	rope := Rope{
		Head: Point{4, 0},
		Tail: Point{3, 0}}
	move := Move{Direction: UP, Times: 4}
	rope.Apply(move)
	assert.Equal(t, Point{4, 4}, rope.Head)
	assert.Equal(t, Point{4, 3}, rope.Tail)
}

func TestRope_ApplyMovePullsRopeBehind_RIGHT_4(t *testing.T) {
	rope := Rope{
		Head: Point{0, 0},
		Tail: Point{0, 0}}
	move := Move{Direction: RIGHT, Times: 4}
	rope.Apply(move)
	assert.Equal(t, Point{4, 0}, rope.Head)
	assert.Equal(t, Point{3, 0}, rope.Tail)
}

func TestRope_CorrectTail_UP_and_align_y_axis(t *testing.T) {
	rope := Rope{
		Head: Point{4, 2},
		Tail: Point{3, 0}}
	expected := Point{4, 1}
	rope.CorrectTail()
	assert.Equal(t, expected, rope.Tail)
}

func TestRope_CorrectTail_DOWN(t *testing.T) {
	rope := Rope{
		Head: Point{3, 0},
		Tail: Point{7, 0}}
	expected := Point{4, 0}
	rope.CorrectTail()
	assert.Equal(t, expected, rope.Tail)
}

func TestRope_CorrectTail_UP(t *testing.T) {
	rope := Rope{
		Head: Point{8, 0},
		Tail: Point{3, 0}}
	expected := Point{7, 0}
	rope.CorrectTail()
	assert.Equal(t, expected, rope.Tail)
}

func TestRope_ApplyMoveDown(t *testing.T) {
	rope := Rope{Head: Point{X: 2, Y: 5}}
	move := Move{Direction: DOWN, Times: 2}
	rope.Apply(move)
	assert.Equal(t, 3, rope.Head.Y)
}

func TestRope_ApplyMoveUp(t *testing.T) {
	rope := Rope{Head: Point{X: 2, Y: 5}}
	move := Move{Direction: UP, Times: 2}
	rope.Apply(move)
	assert.Equal(t, 7, rope.Head.Y)
}

func TestRope_ApplyMoveRight(t *testing.T) {
	rope := Rope{Head: Point{X: 2, Y: 5}}
	move := Move{Direction: RIGHT, Times: 2}
	rope.Apply(move)
	assert.Equal(t, 4, rope.Head.X)
}

func TestRope_ApplyMoveLeft(t *testing.T) {
	rope := Rope{Head: Point{X: 2, Y: 5}}
	move := Move{Direction: LEFT, Times: 2}
	rope.Apply(move)
	assert.Equal(t, 0, rope.Head.X)
}

func TestParseInput(t *testing.T) {
	expected := []Move{
		{Direction: UP, Times: 1},
		{Direction: LEFT, Times: 3},
	}
	input := "U 1\nL 3\n"
	moves := parseInput(strings.NewReader(input))
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
