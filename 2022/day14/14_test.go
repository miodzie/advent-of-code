package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

var example = "498,4 -> 498,6 -> 496,6\n503,4 -> 502,4 -> 502,9 -> 494,9"

func TestSolutionTwo(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	graph := make(Graph)
	graph.Plot(ParseInput(f))
	assert.Equal(t, 29076, part2(Coord{500, 0}, graph))
}

func TestExampleTwo(t *testing.T) {
	graph := getExample()
	assert.Equal(t, 93, part2(Coord{500, 0}, graph))
}

func TestSolutionOne(t *testing.T) {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	graph := make(Graph)
	graph.Plot(ParseInput(f))
	assert.Equal(t, 799, part1(Coord{500, 0}, graph))
}

func TestExampleOne(t *testing.T) {
	graph := getExample()
	assert.Equal(t, 24, part1(Coord{500, 0}, graph))
}

func TestDropSand(t *testing.T) {
	graph := getExample()
	sand := Coord{500, 0}

	graph.DropSand(sand)

	assert.True(t, graph[Coord{500, 8}])
	assert.Len(t, graph, 21)
	graph.DropSand(sand)
	assert.True(t, graph[Coord{499, 8}])
	assert.Len(t, graph, 22)
}

func TestLine_Plot_X(t *testing.T) {
	line := Line{A: Coord{X: 498, Y: 4}, B: Coord{X: 498, Y: 6}}
	graph := make(Graph)
	line.Plot(graph)
	assert.Len(t, graph, 3)
	assert.True(t, graph[line.A])
	assert.True(t, graph[Coord{X: 498, Y: 5}])
	assert.True(t, graph[line.B])
}

func TestParse(t *testing.T) {
	r := strings.NewReader(example)
	lines := ParseInput(r)
	line1 := Line{Coord{X: 498, Y: 4}, Coord{X: 498, Y: 6}}
	line2 := Line{Coord{X: 498, Y: 6}, Coord{X: 496, Y: 6}}
	assert.NotNil(t, lines)
	assert.Equal(t, lines[0], line1)
	assert.Equal(t, lines[1], line2)
}

func TestPlotsExample(t *testing.T) {
	assert.Len(t, getExample(), 20)
}

////////////////////////////////////////////////////////////////////////////////////

func getExample() (graph Graph) {
	graph = make(Graph)
	graph.Plot(ParseInput(strings.NewReader(example)))
	return graph
}
