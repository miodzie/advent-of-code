package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		panic(err)
	}
	input := ParseInput(f)
	graph := ParseGraph(input)
	Q := queue[Node]{}
	for _, n := range graph {
		if n.val == 'S' {
			Q.Push(n)
			break
		}
	}
	Q.items[0].dump()
	visited := make(map[int]bool)
	node := Q.Shift()
	for ; node.val != 'E'; node = Q.Shift() {
		if visited[node.id] {
			continue
		}
		visited[node.id] = true
		// Find the smallest cost,
		// Add it to the queue.
		max, next := math.MinInt, &Node{}
		for _, n := range node.neighbors {
			n.distance = node.distance + 1
			Q.Push(n)
			cost := n.cost() - node.cost()
			////fmt.Printf("Cost of %q: %d, cost: %d\n", n.val, n.cost(), cost)
			if cost >= max && !visited[n.id] {
				max = cost
				next = n
			}
		}
		if next.id == 0 {
			continue
		}
		//print("Pushing: ")
		//next.dump()
		Q.Push(next)
	}
	fmt.Println(node.distance)
}

func ParseGraph(input [][]*Node) (graph []*Node) {
	for row := range input {
		for col := range input[row] {
			node := input[row][col]
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
	id        int
	val       rune
	neighbors []*Node
	_weight   int
	distance  int
}

func (n *Node) weight() int {
	return n.cost() + n._weight
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
	fmt.Printf("%s: neighbors: %s\n", string(n.val), strings.Join(neigh, ","))
}

func (n *Node) add(node *Node) {
	n.neighbors = append(n.neighbors, node)
}

type queue[T any] struct {
	items []*T
}

func (q *queue[T]) Shift() *T {
	var p *T
	p, q.items = q.items[0], q.items[1:]
	return p
}

func (q *queue[T]) Pop() *T {
	var p *T
	p, q.items = q.items[len(q.items)-1], q.items[:len(q.items)-1]
	return p
}

func (q *queue[T]) Push(node *T) {
	q.items = append(q.items, node)
}

func (q *queue[T]) Empty() bool {
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
