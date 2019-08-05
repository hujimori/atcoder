package main

import (
	"bufio"
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
	to int
	c  int
	t  int
}

var G [][]edge
var N, M int

const INF = 1001000000000.0

func main() {
	sc.Split(bufio.ScanWords)
	N, M = nextInt(), nextInt()

	G = make([][]edge, N)

	for i := 0; i < M; i++ {
		a, b, c, d := nextInt(), nextInt(), nextInt(), nextInt()
		G[a] = append(G[a], edge{to: b, c: c, t: d})
		G[b] = append(G[b], edge{to: a, c: c, t: d})

	}

	lb := 0.0
	ub := INF

	for i := 0; i < 10000; i++ {
		mid := (lb + ub) / 2

		if C(mid) {
			lb = mid
		} else {
			ub = mid
		}
	}

}

func C(x float64) bool {

	for i := 0; i < N; i++ {
		seen := make([]int, N)
		dfs
	}

}

func dfs() bool {

}
