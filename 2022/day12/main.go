package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strings"
)

var a = &Node{distance: math.MaxInt}

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	input := ParseInput(f)
	graph := ParseGraph(input)
	var start *Node
	for _, n := range graph {
		if n.val == 'S' {
			start = n
			break
		}
	}
	fmt.Println(shortestPath(start))
	reset(graph)
	smallest := []int{}
	for _, n := range graph {
		if n.val == 'a' {
			reset(graph)
			steps := shortestPath(n)
			if steps > 0 {
				smallest = append(smallest, steps)
			}
		}
	}
	// 9223372036854775807 is too high.
	fmt.Println(a.distance)
	sort.Ints(smallest)
	// 381 is too high
	fmt.Println(smallest[0])
}

func reset(graph []*Node) {
	for _, n := range graph {
		n.reset()
	}
}

func shortestPath(start *Node) int {
	Q := queue{}
	Q.Push(start)
	node := start
	for node.val != 'E' {
		if Q.Empty() {
			return -1
		}
		node = Q.Shift()
		if node.val == 'E' {
		}
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
			if Q.Empty() {
				break
			}
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
	neigh := []string{}
	for _, neighbor := range n.neighbors {
		neigh = append(neigh, string(neighbor.val))
	}
	fmt.Printf("%s: coord: %s neighbors: %s\n", string(n.val), n._id,
		strings.Join(neigh, ","))
}

func (n *Node) add(node *Node) {
	n.neighbors = append(n.neighbors, node)
}

type queue struct {
	items []*Node
}

func (q *queue) Shift() *Node {
	var p *Node
	if len(q.items) == 1 {
		//q.items[0].dump()
	}
	p, q.items = q.items[0], q.items[1:]
	return p
}

func (q *queue) Pop() *Node {
	var p *Node
	p, q.items = q.items[len(q.items)-1], q.items[:len(q.items)-1]
	return p
}

func (q *queue) Push(node *Node) {
	q.items = append(q.items, node)
}

func (q *queue) NotEmpty() bool {
	return len(q.items) != 0
}

func (q *queue) Empty() bool {
	return len(q.items) == 0
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
