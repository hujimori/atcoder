package main

import (
	"bufio"
	"container/heap"
	"fmt"
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

type edge struct {
	to   int
	cost int
}

type pair struct {
	first  int
	second int
}

type PriorityQueue []*pair

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].first < pq[j].first
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	pair := old[n-1]
	*pq = old[0 : n-1]
	return pair
}

func (pq *PriorityQueue) Push(x interface{}) {
	pair := x.(*pair)
	*pq = append(*pq, pair)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

const MAX = 100100
const INF = 100000000000

func dijkstra(s int, G [][]edge) []int {
	pq := make(PriorityQueue, N)
	for i := 0; i < N; i++ {
		pq[i] = &pair{}
	}
	heap.Init(&pq)

	d := make([]int, MAX)
	for i := 0; i < MAX; i++ {
		d[i] = INF
	}
	d[s] = 0
	heap.Push(&pq, &pair{first: 0, second: s})

	for pq.Len() > 0 {
		p := heap.Pop(&pq).(*pair)
		v := p.second
		if d[v] < p.first {
			continue
		}
		for _, e := range G[v] {
			if d[e.to] > d[v]+e.cost {
				d[e.to] = d[v] + e.cost
				heap.Push(&pq, &pair{first: d[e.to], second: e.to})
			}
		}
	}
	return d
}

var G [][]edge
var rG [][]edge
var N, M, T int
var A []int

func main() {
	sc.Split(bufio.ScanWords)

	N, M, T = nextInt(), nextInt(), nextInt()
	A = make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = nextInt()
	}
	G = make([][]edge, N)
	rG = make([][]edge, N)
	for i := 0; i < M; i++ {
		a, b, c := nextInt()-1, nextInt()-1, nextInt()
		G[a] = append(G[a], edge{to: b, cost: c})
		rG[b] = append(rG[b], edge{to: a, cost: c})
	}

	d := dijkstra(0, G)
	rd := dijkstra(0, rG)
	ans := 0
	for i := 0; i < N; i++ {
		t := d[i] + rd[i]
		if T-t >= 0 {
			ans = max(ans, (T-t)*A[i])
		}
	}

	fmt.Println(ans)
}

func max(a, b int) int {
	if a < b {
		return b

	}

	return a
}
