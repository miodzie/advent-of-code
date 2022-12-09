package main

import (
	"bufio"
	"fmt"
	"strings"
)

const (
	UP    = 'U'
	LEFT  = 'L'
	DOWN  = 'D'
	RIGHT = 'R'
)

type Move struct {
	Direction rune
	Times     int
}

type Point struct{ X, Y int }

type Rope struct {
	Head, Tail     Point
	PrevTailPoints []Point
}

func (r *Rope) Apply(move Move) {
	switch move.Direction {
	case RIGHT:
		r.Head.X += move.Times
		if move.Times > 1 {
			r.Tail = r.Head
			r.Tail.X -= 1
		}
		break
	case LEFT:
		r.Head.X -= move.Times
		if move.Times > 1 {
			r.Tail = r.Head
			r.Tail.X += 1
		}
		break
	case UP:
		r.Head.Y += move.Times
		if move.Times > 1 {
			r.Tail = r.Head
			r.Tail.X -= 1
		}
	case DOWN:
		r.Head.Y -= move.Times
		if move.Times > 1 {
			r.Tail = r.Head
			r.Tail.Y += 1
		}
	}
}

func (r *Rope) CorrectTail() {
	if r.Head.X-1 > r.Tail.X {
		r.Tail.X += (r.Head.X - 1) - r.Tail.X
	}
	if r.Head.X+1 < r.Tail.X {
		r.Tail.X -= r.Tail.X - (r.Head.X + 1)
	}

	if r.Head.Y-1 > r.Tail.Y {
		r.Tail.Y += (r.Head.Y - 1) - r.Tail.Y
		// Was a diagonal correction, line them up.
		if r.Head.X > r.Tail.X {
			r.Tail.X = r.Head.X
		}
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
