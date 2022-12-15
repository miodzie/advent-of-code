package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

var maxY int

func part2(sandStart Coord, graph Graph) (sum int) {
	maxY = graph.MaxY() + 2
	(Line{Coord{-200_000, maxY},
		Coord{X: 200_000, Y: maxY}}).Plot(graph)
	for !graph[sandStart] {
		graph.DropSand(sandStart)
		sum++
	}
	return sum
}

func part1(sandStart Coord, graph Graph) (sum int) {
	maxY = graph.MaxY()
	for graph.DropSand(sandStart) {
		sum++
	}
	return sum
}

func (g Graph) DropSand(sand Coord) bool {
	// We went past the lowest line, or already exists.
	// TODO: globals like maxY are just terrible but w/e
	if sand.Y >= maxY || g[sand] {
		return false
	}
	unitBelow := Coord{X: sand.X, Y: sand.Y + 1}
	// if there's a unitBelow.
	if g[unitBelow] {
		l := Coord{unitBelow.X - 1, unitBelow.Y}
		// if there's a spot on the left, try that.
		if !g[l] {
			return g.DropSand(l)
		}
		r := Coord{unitBelow.X + 1, unitBelow.Y}
		// if there's a spot on the right, try that.
		if !g[r] {
			return g.DropSand(r)
		}
		// if there's a unit on the left and right,
		// Sand has nowhere else to go.
		if g[l] && g[r] {
			g[sand] = true
			return true
		}
	}
	sand.Y++
	return g.DropSand(sand)
}

type Graph map[Coord]bool

func (g Graph) Plot(lines []Line) {
	for _, line := range lines {
		line.Plot(g)
	}
}

func (g Graph) MaxY() (Y int) {
	for l, _ := range g {
		if l.Y >= Y {
			Y = l.Y
		}
	}
	return Y
}

type Line struct{ A, B Coord }

func (l Line) Plot(graph map[Coord]bool) {
	graph[l.A] = true
	graph[l.B] = true

	diff := l.Diff()
	if diff.X != 0 {
		var min int
		var max int
		if l.A.X < l.B.X {
			min, max = l.A.X, l.B.X
		} else {
			min, max = l.B.X, l.A.X
		}
		for x := min; x < max; x++ {
			graph[Coord{X: x, Y: l.A.Y}] = true
		}
	}
	if diff.Y != 0 {
		var min int
		var max int
		if l.A.Y < l.B.Y {
			min, max = l.A.Y, l.B.Y
		} else {
			min, max = l.B.Y, l.A.Y
		}
		for y := min; y < max; y++ {
			graph[Coord{X: l.A.X, Y: y}] = true
		}
	}
}

func (l Line) Diff() Coord {
	return Coord{X: l.A.X - l.B.X, Y: l.A.Y - l.B.Y}
}

type Coord struct {
	X, Y int
}

func ParseInput(r io.Reader) (lines []Line) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		cords := strings.Split(line, "->")
		var l Line
		for _, c := range cords {
			var p Coord
			_, _ = fmt.Sscanf(c, "%d,%d", &p.X, &p.Y)
			if l.A == (Coord{}) {
				l.A = p
			} else {
				l.B = p
				lines = append(lines, l)
				l = Line{A: p}
			}
		}
	}
	return lines
}
