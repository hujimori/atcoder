package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

type Item struct {
	First  int // 最短距離
	Second int // 頂点の番号
	Index  int
}

type PriorityQueue []*Item

func NewPriorityQueue(n int) PriorityQueue {
	p := make(PriorityQueue, n)
	for i := 0; i < n; i++ {
		p[i] = &Item{}
		p[i].Index = i
	}
	heap.Init(&p)
	return p
}

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].First < pq[j].First
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = j
}

var priorityQueue PriorityQueue

type edge struct {
	to   int
	cost int
}

var g [][]edge
var (
	dist  []float64 // 最短距離
	dist2 []float64 // 2番目の最短距離
)
var (
	N int
	R int
)

func main() {
	sc.Split(bufio.ScanWords)
	N, R = nextInt(), nextInt()
	g = make([][]edge, N)
	for i := 0; i < N; i++ {
		g[i] = make([]edge, 0)
	}
	for i := 0; i < N; i++ {
		u, v, w := nextInt()-1, nextInt()-1, nextInt()
		g[u] = append(g[u], edge{v, w})
		g[v] = append(g[v], edge{u, w})
	}

	fmt.Println(g)
	dijkstra()
	// a := 0.0
	// b := 1.0
	// fmt.Printf("a=%f b=%f\n", a, b)
	// swap(&a, &b)
	// fmt.Printf("a=%f b=%f\n", a, b)

}

func dijkstra() {
	p := NewPriorityQueue(N)
	dist = make([]float64, N)
	dist2 = make([]float64, N)
	for i := 0; i < N; i++ {
		dist[i] = math.Inf(1)
		dist2[i] = math.Inf(1)
	}

	dist[0] = 0
	heap.Push(&p, &Item{First: 0, Second: 0})

	for p.Len() != 0 {
		i := heap.Pop(&p).(*Item)
		// fmt.Println(i)
		d := i.First  // 最短距離
		v := i.Second // 頂点の番号

		if dist2[v] < float64(d) {
			continue
		}

		for i := 0; i < len(g[v]); i++ {
			e := g[v][i]
			d2 := float64(d + e.cost)
			// fmt.Println(d2)

			if dist[e.to] > d2 {

				swap(&dist[e.to], &d2)
				heap.Push(&p, &Item{First: int(dist[e.to]), Second: int(e.to)})
			}

			if dist2[e.to] > d2 && dist[e.to] < d2 {
				dist2[e.to] = d2
				heap.Push(&p, &Item{First: int(dist2[e.to]), Second: int(e.to)})
			}

		}

	}

	println(int(dist2[N-1]))
}

func swap(a, b *float64) {
	tmp := *a
	*a = *b
	*b = tmp
}
