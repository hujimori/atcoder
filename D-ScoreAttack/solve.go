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

var N, M int
var G [][]edge

const MAX = 1100
const INF = 110000000000000

func main() {
	sc.Split(bufio.ScanWords)

	N, M = nextInt(), nextInt()
	G = make([][]edge, N)

	for i := 0; i < M; i++ {
		u, v, c := nextInt()-1, nextInt()-1, nextInt()
		G[u] = append(G[u], edge{to: v, cost: -c})
	}

	d := make([]int, MAX)
	for i := 0; i < MAX; i++ {
		d[i] = INF
	}
	negative := false
	d[0] = 0

	for iter := 0; iter <= N*2; iter++ {
		for v := 0; v < N; v++ {
			if d[v] >= INF/2 {
				continue
			}
			for _, e := range G[v] {

				if d[e.to] > d[v]+e.cost {
					d[e.to] = d[v] + e.cost
					if e.to == N-1 && iter == N*2 {
						negative = true
					}
				}
			}
		}
	}

	if !negative {
		fmt.Println(-d[N-1])
	} else {
		fmt.Println("inf")
	}

}
