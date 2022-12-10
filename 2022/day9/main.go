package main

import (
	"bufio"
	"fmt"
	"github.com/miodzie/advent-of-code/util"
	"io"
)

func part2(moves []Move) int {
	head := NewKnot(0)
	tail := head
	for i := 1; i < 10; i++ {
		tail.append(NewKnot(i))
		tail = tail.next
	}

	for _, m := range moves {
		head.move(m)
	}

	return len(tail.visited)
}

func part1(moves []Move) int {
	head := NewKnot(0)
	head.append(NewKnot(1))

	for _, m := range moves {
		head.move(m)
	}

	return len(head.next.visited)
}

type coordinates struct{ X, Y int }

type Knot struct {
	id      int
	coords  coordinates
	visited map[coordinates]bool
	next    *Knot
}

func NewKnot(id int) *Knot {
	return &Knot{id: id, visited: make(map[coordinates]bool)}
}

func (k *Knot) append(knot *Knot) {
	k.next = knot
}

func (k *Knot) touching(knot *Knot) bool {
	x := util.Abs(knot.coords.X - k.coords.X)
	y := util.Abs(knot.coords.Y - k.coords.Y)

	return 1 >= x && 1 >= y
}

func (k *Knot) follow(knot *Knot) {
	if k.touching(knot) {
		if k.next == nil {
			k.visited[k.coords] = true
			return
		} else {
			k.next.follow(k)
			return
		}
	}

	diffX := knot.coords.X - k.coords.X
	diffY := knot.coords.Y - k.coords.Y

	// diagonal
	if diffX != 0 && diffY != 0 {
		if util.Abs(diffY) > 0 {
			if diffY < 0 {
				k.coords.Y--
			} else {
				k.coords.Y++
			}
		}
		if util.Abs(diffX) > 0 {
			if diffX < 0 {
				k.coords.X--
			} else {
				k.coords.X++
			}
		}
	} else {
		if util.Abs(diffY) > 1 {
			if diffY < 0 {
				k.coords.Y--
			} else {
				k.coords.Y++
			}
		}

		if util.Abs(diffX) > 1 {
			if diffX < 0 {
				k.coords.X--
			} else {
				k.coords.X++
			}
		}
	}

	k.follow(knot)
}

func (k *Knot) move(move Move) {
	incr := 1
	dir := move.Direction
	if dir == 'L' || dir == 'D' {
		incr *= -1
	}
	for i := 0; i < move.Times; i++ {
		if dir == 'U' || dir == 'D' {
			k.coords.Y += incr
		}
		if dir == 'L' || dir == 'R' {
			k.coords.X += incr
		}
		if k.next != nil {
			k.next.follow(k)
		}
	}
}

type Move struct {
	Direction rune
	Times     int
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
