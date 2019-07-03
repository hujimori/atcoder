package main

import "fmt"

type Edge struct {
	to   int
	cost int
}

type Graph []*Edge

func NewGraph(n int) Graph {
	g := make(Graph, n)
	return g
}

func (g Graph) PushBack(i int, e *Edge) {
	g[i] = e
}

func main() {

	g := NewGraph(3)
	g.PushBack(0, &Edge{to: 1, cost: 10})
	g.PushBack(1, &Edge{to: 2, cost: 5})
	g.PushBack(2, &Edge{to: 0, cost: 3})

	fmt.Println(g[0])
	fmt.Println(g[1])
	fmt.Println(g[2])
}
