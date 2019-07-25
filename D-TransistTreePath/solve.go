package main

import (
	"bufio"
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

var G [][]edge
var N int
var a, b, c int
var Q, K int
var x, y []int
var dist []int

func main() {
	sc.Split(bufio.ScanWords)

	N = nextInt()
	G = make([][]edge, N)
	for i := 0; i < N-1; i++ {
		a, b, c = nextInt()-1, nextInt()-1, nextInt()
		G[a] = append(G[a], edge{to: b, cost: c})
		G[b] = append(G[b], edge{to: a, cost: c})
	}

	Q, K = nextInt(), nextInt()-1
	x = make([]int, Q)
	y = make([]int, Q)
	for q := 0; q < Q; q++ {
		x[q], y[q] = nextInt()-1, nextInt()-1
	}
	dist = make([]int, N)
	rec(K, -1, 0)
	for q := 0; q < Q; q++ {
		fmt.Println(dist[x[q]] + dist[y[q]])
	}
}

func rec(v, p, sum int) {
	dist[v] = sum
	for _, e := range G[v] {
		if e.to == p {
			continue
		}
		rec(e.to, v, sum+e.cost)

	}
}
