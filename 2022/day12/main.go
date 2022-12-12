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
	f, err := os.Open("example")
	if err != nil {
		panic(err)
	}
	input := ParseInput(f)
	graph := ParseGraph(input)
	Q := queue[Node]{}
	graph[0].dump()
	Q.Push(graph[0]) // S is 0,0 in example.
	seen := make(map[int]bool)
	for node := Q.Pop(); node.val != 'E'; node = Q.Pop() {
		seen[node.id] = true
		if node.val == 'E' {
			fmt.Println("MY PANTS")
			fmt.Println(node.val)
			return
		}
		// Find the smallest cost,
		// Add it to the queue.
		max, next := math.MinInt, &Node{}
		for _, n := range node.neighbors {
			if n.val == 'S' {
				continue
			}
			cost := n.cost() - node.cost()
			if n.cost()-node.cost() > 1 {
				continue
			}
			//fmt.Printf("Cost of %q: %d, cost: %d\n", n.val, n.cost(), cost)
			if cost >= max && !seen[n.id] {
				max = cost
				next = n
			}
		}
		next.distance += node.distance + 1
		if next.val == 'E' {
			next.dump()
			fmt.Println(next.distance)
		}
		if next.id == 0 {
			panic("I couldn't find a neighbor :(")
		}
		print("Pushing: ")
		next.dump()
		//time.Sleep(1 * time.Second)
		Q.Push(next)
	}
}

func ParseGraph(input [][]*Node) (graph []*Node) {
	for y := range input {
		for x := range input[y] {
			node := input[y][x]
			// gather neighbors
			// TODO: Can replace this with a fancy for loop.
			// up
			if y > 0 {
				node.add(input[y-1][x])
			}
			// down
			if y < len(input)-1 {
				node.add(input[y+1][x])
			}
			// left
			if x > 0 {
				node.add(input[y][x-1])
			}
			// right
			if x < len(input[y])-1 {
				node.add(input[y][x+1])
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
