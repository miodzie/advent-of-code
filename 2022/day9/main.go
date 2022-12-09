package main

import (
	"bufio"
	"fmt"
	"strings"
)

type Direction rune

const (
	UP    Direction = 'U'
	LEFT  Direction = 'L'
	DOWN  Direction = 'D'
	RIGHT Direction = 'R'
)

type Move struct {
	Direction Direction
	Times     int
}

func (m Move) Transformation() Point {
	switch m.Direction {
	case UP:
		return Point{0, m.Times}
	case DOWN:
		return Point{0, -m.Times}
	case LEFT:
		return Point{-m.Times, 0}
	case RIGHT:
		return Point{m.Times, 0}
	}
	return Point{}
}

type Point struct{ X, Y int }

func (p *Point) Diff(two Point) Point {
	return Point{p.X - two.X, p.Y - two.Y}
}

func (p *Point) Transform(t Point) {
	p.X += t.X
	p.Y += t.Y
}

func (p *Point) Equals(t Point) bool {
	return p.X == t.X && p.Y == t.Y
}

type Rope struct {
	Head, Tail     Point
	PrevTailPoints []Point
}

func (r *Rope) ApplyMove(move Move) {
	transformation := move.Transformation()
	oldHead := r.Head
	preDiff := r.Head.Diff(r.Tail)

	r.Head.Transform(transformation)

	tpm := move
	tpm.Times -= 1
	r.Tail.Transform(move.Transformation())

	// Realign on the X axis
	if preDiff.X >= 1 && move.Times > 1 {
		r.Tail.X = r.Head.X
	}
	// Realign on Y axis
	if preDiff.Y >= 1 && move.Times > 1 {
		r.Tail.Y = r.Head.Y
	}
	// Realign if diagonal.
	if preDiff.Equals(Point{1, 1}) && move.Times == 1 {
		r.Tail = oldHead
	}
}

func parseInput(r *strings.Reader) (moves []Move) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		moves = append(moves, parseMove(scanner.Text()))
	}
	return
}

func parseMove(s string) (move Move) {
	fmt.Sscanf(s, "%c %d", &move.Direction, &move.Times)
	return
}
