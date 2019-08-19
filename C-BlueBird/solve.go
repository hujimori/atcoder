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

var G []edge
var N int
var M int
var d [310][310]int

func warshallFloyd() {
	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				d[i][j] = min(d[i][j], d[i][k]+d[j][k])
			}
		}
	}
}

const INF = 1000000000000

func main() {
	sc.Split(bufio.ScanWords)

	for i := 0; i < 310; i++ {
		for j := 0; j < 310; j++ {
			d[i][j] = INF
		}
	}

	N, M = nextInt(), nextInt()
	G = make([]edge, N)
	for i := 0; i < M; i++ {
		u, v, l := nextInt()-1, nextInt()-1, nextInt()

		if u == 0 {
			G = append(G, edge{cost: l, to: v})
			continue
		} else if v == 0 {
			G = append(G, edge{cost: l, to: u})
			continue
		}

		d[u][v] = l
		d[v][u] = l
	}

	warshallFloyd()
	ans := INF
	for _, u := range G {
		for _, v := range G {
			if u.to == v.to {
				continue
			}
			l := u.cost + d[u.to][v.to] + v.cost
			ans = min(ans, l)
		}
	}

	if ans == INF {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
