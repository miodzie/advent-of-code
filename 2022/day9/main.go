package main

import (
	"bufio"
	"fmt"
	"github.com/miodzie/advent-of-code/util"
	"io"
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

type Point struct{ X, Y int }

func (p Point) String() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

type Rope struct {
	Head, Tail Point
}

func AllMoves(rope *Rope, moves []Move) int {
	seenTails := make(map[string]any)

	for _, m := range moves {
		tails := applyMove(rope, m)
		for _, t := range tails {
			seenTails[t.String()] = 1
		}
	}

	return len(seenTails)
}

func applyMove(rope *Rope, move Move) []Point {
	var tailPoints = []Point{rope.Tail}

	for i := move.Times; i > 0; i-- {
		prevHead := rope.Head
		switch move.Direction {
		case UP:
			rope.Head.Y += 1
			break
		case LEFT:
			rope.Head.X -= 1
			break
		case DOWN:
			rope.Head.Y -= 1
			break
		case RIGHT:
			rope.Head.X += 1
			break
		}
		x := util.Abs(rope.Head.X - rope.Tail.X)
		y := util.Abs(rope.Head.Y - rope.Tail.Y)
		//fmt.Printf("%d, %d\n", x, y)
		if x > 1 || y > 1 {
			rope.Tail = prevHead
			tailPoints = append(tailPoints, rope.Tail)
		}
	}

	return tailPoints
}

func ParseInput(r io.Reader) (moves []Move) {
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
