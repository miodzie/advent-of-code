package main

import (
	"bufio"
	"fmt"
	"github.com/miodzie/advent-of-code/util"
	"io"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	input := ParseInput(f)
	graph := ParseGraph(input)
	var start *Node
	var destination *Node
	for _, n := range graph {
		if n.val == 'S' {
			start = n
		}
		if n.val == 'E' {
			destination = n
		}
	}
	fmt.Println(search(start, destination))
	reset(graph)
	var smallest []int
	for _, n := range graph {
		if n.val == 'a' {
			reset(graph)
			steps := search(n, destination)
			if steps > 0 {
				smallest = append(smallest, steps)
			}
		}
	}
	sort.Ints(smallest)
	fmt.Println(smallest[0])
}

func reset(graph []*Node) {
	for _, n := range graph {
		n.reset()
	}
}

func search(start *Node, destination *Node) int {
	Q := util.Queue[Node]{}
	Q.Push(start)
	node := start
	for node != destination {
		if Q.Empty() {
			return -1
		}
		node = Q.Shift()
		if node.visited {
			continue
		}
		node.visited = true
		max, next := math.MinInt, &Node{}
		for _, n := range node.neighbors {
			n.distance = node.distance + 1
			Q.Push(n)
			cost := n.cost() - node.cost()
			if cost >= max && !n.visited {
				max = cost
				next = n
			}
		}
		if next.id == 0 {
			continue
		}
		Q.Push(next)
	}

	return node.distance
}

func ParseGraph(input [][]*Node) (graph []*Node) {
	for row := range input {
		for col := range input[row] {
			node := input[row][col]
			node._id = fmt.Sprintf("%d,%d", row, col)
			// gather neighbors
			for _, set := range [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}} {
				rr, cc := row+set[0], col+set[1]
				if (0 <= rr && rr <= len(input)-1) &&
					0 <= cc && cc <= len(input[row])-1 {
					neighbor := input[rr][cc]
					cost := neighbor.cost() - node.cost()
					// Only step-able neighbors
					if cost < 2 {
						node.add(neighbor)
					}
				}
			}
			graph = append(graph, node)
		}
	}

	return
}

type Node struct {
	val       rune
	id        int
	neighbors []*Node
	visited   bool
	distance  int
	_id       string
}

func (n *Node) reset() {
	n.distance = 0
	n.visited = false
}

func (n *Node) cost() int {
	if n.val == 'S' {
		return 0
	}
	if n.val == 'E' {
		return 27
	}
	return int(n.val) - 96
}

func (n *Node) dump() {
	var neigh []string
	for _, neighbor := range n.neighbors {
		neigh = append(neigh, string(neighbor.val))
	}
	fmt.Printf(
		"%s: coord: %s neighbors: %s\n",
		string(n.val), n._id, strings.Join(neigh, ","))
}

func (n *Node) add(node *Node) {
	n.neighbors = append(n.neighbors, node)
}

func ParseInput(reader io.Reader) (graph [][]*Node) {
	scanner := bufio.NewScanner(reader)
	n := 0
	for scanner.Scan() {
		var edges []*Node
		line := strings.TrimSpace(scanner.Text())
		for _, c := range line {
			edges = append(edges, &Node{val: c, id: n})
			n++
		}
		graph = append(graph, edges)
	}
	return graph
}
