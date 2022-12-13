package util

// TODO: Write a generic BFS, Dijkstra? Seems like it could be useful.

type Queue[T any] struct {
	items []*T
}

func (q *Queue[T]) Shift() *T {
	var p *T
	p, q.items = q.items[0], q.items[1:]
	return p
}

func (q *Queue[T]) Pop() *T {
	var p *T
	p, q.items = q.items[len(q.items)-1], q.items[:len(q.items)-1]
	return p
}

func (q *Queue[T]) Push(node *T) {
	q.items = append(q.items, node)
}

func (q *Queue[T]) Empty() bool {
	return len(q.items) == 0
}
