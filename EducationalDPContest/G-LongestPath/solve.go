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
	node int
}

var g [][]edge
var n int
var m int
var dp [100100]float64

func rec(v int) int {
	if dp[v] != -1 {
		return int(dp[v])
	}

	res := 0.0
	for i := 0; i < len(g[v]); i++ {
		chmax(&res, float64(rec(g[v][i].node))+1)
	}
	dp[v] = res
	return int(dp[v])
}

func main() {
	sc.Split(bufio.ScanWords)
	n = nextInt()
	m = nextInt()
	g = make([][]edge, n)
	for i := 0; i < n; i++ {
		g[i] = make([]edge, 0)
	}

	for i := 0; i < m; i++ {
		u, v := nextInt()-1, nextInt()-1
		g[u] = append(g[u], edge{v})
	}

	// 初期化
	for i := 0; i < n; i++ {
		dp[i] = -1
	}

	res := 0.0
	for i := 0; i < n; i++ {
		chmax(&res, float64(rec(i)))
	}
	fmt.Println(res)

}

func chmin(a *float64, b float64) bool {
	if *a > b {
		*a = b
		return true
	}
	return false
}

func chmax(a *float64, b float64) bool {
	if *a < b {
		*a = b
		return true
	}
	return false
}
