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

func AllMoves(rope []Point, moves []Move) int {
	seenTails := make(map[string]any)

	for _, m := range moves {
		tails := applyMove(rope, m)
		for _, t := range tails {
			seenTails[t.String()] = 1
		}
	}

	return len(seenTails)
}

func applyMove(rope []Point, move Move) []Point {
	head := rope[0]
	tail := rope[len(rope)-1]
	var tailPoints = []Point{tail}

	for i := move.Times; i > 0; i-- {
		prevHead := head
		switch move.Direction {
		case UP:
			head.Y += 1
			break
		case LEFT:
			head.X -= 1
			break
		case DOWN:
			head.Y -= 1
			break
		case RIGHT:
			head.X += 1
			break
		}
		x := util.Abs(head.X - tail.X)
		y := util.Abs(head.Y - tail.Y)
		//fmt.Printf("%d, %d\n", x, y)
		if x > 1 || y > 1 {
			tail = prevHead
			tailPoints = append(tailPoints, tail)
		}
		rope[0], rope[1] = head, tail
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
