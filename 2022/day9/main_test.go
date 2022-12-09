package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestRope_ApplyMove_logs_previous_tail_positions(t *testing.T) {

}

func TestRope_ApplyMove_realigns_if_was_diagonal(t *testing.T) {
	rope := Rope{Head: Point{3, 3}, Tail: Point{2, 2}}
	move := Move{UP, 1}
	rope.ApplyMove(move)
	assert.Equal(t, Point{3, 4}, rope.Head)
	assert.Equal(t, Point{3, 3}, rope.Tail)
}

func TestRope_ApplyMove_LEFT_3_realigns_Y_axis(t *testing.T) {
	rope := Rope{Head: Point{4, 4}, Tail: Point{4, 3}}
	move := Move{LEFT, 3}
	rope.ApplyMove(move)
	assert.Equal(t, Point{1, 4}, rope.Head)
	assert.Equal(t, Point{2, 4}, rope.Tail)
}

func TestRope_ApplyMove_UP_4_realigns_X_axis(t *testing.T) {
	rope := Rope{Head: Point{4, 0}, Tail: Point{3, 0}}
	move := Move{UP, 4}
	rope.ApplyMove(move)
	assert.Equal(t, Point{4, 4}, rope.Head)
	assert.Equal(t, Point{4, 3}, rope.Tail)
}

func TestRope_ApplyMove(t *testing.T) {
	rope := Rope{Head: Point{}, Tail: Point{}}
	move := Move{RIGHT, 4}

	rope.ApplyMove(move)

	assert.Equal(t, Point{4, 0}, rope.Head)
	assert.Equal(t, Point{3, 0}, rope.Tail)
}

func TestMove_CreateTransformation(t *testing.T) {
	up := Move{Direction: UP, Times: 5}
	assert.Equal(t, Point{0, 5}, up.Transformation())
	down := Move{Direction: DOWN, Times: 2}
	assert.Equal(t, Point{0, -2}, down.Transformation())
	left := Move{Direction: LEFT, Times: 2}
	assert.Equal(t, Point{-2, 0}, left.Transformation())
	right := Move{Direction: RIGHT, Times: 2}
	assert.Equal(t, Point{2, 0}, right.Transformation())
}

func TestPoint_Transform_applies_the_given_transformation(t *testing.T) {
	point := Point{4, 5}
	transformation := Point{2, 4}
	point.Transform(transformation)
	assert.Equal(t, Point{6, 9}, point)
}

func TestPoint_Diff_returns_a_point_with_the_differences(t *testing.T) {
	one := Point{3, 7}
	two := Point{1, 3}
	expected := Point{2, 4}
	assert.Equal(t, expected, one.Diff(two))
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
