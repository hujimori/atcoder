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

var colors []int
var g [][]edge
var n int

func main() {
	sc.Split(bufio.ScanWords)
	n = nextInt()
	g = make([][]edge, n)
	for i := 0; i < n; i++ {
		g[i] = make([]edge, 0)
	}
	colors = make([]int, n)

	for i := 0; i < n; i++ {
		u, v := nextInt(), nextInt()
		g[u] = append(g[u], edge{v})
		g[v] = append(g[v], edge{u})
	}

	fmt.Println(g)

	solve()

}

func dfs(v, c int) bool {
	colors[v] = c // 頂点vをCで塗る
	for i := 0; i < len(g[v]); i++ {
		// 隣接している頂点が同じ色ならfalse
		if colors[g[v][i].node] == c {
			return false
		}
		// 隣接している色がまだ塗られていないら-cで塗る
		if colors[g[v][i].node] == 0 && !dfs(g[v][i].node, -c) {
			return false
		}
	}

	// すべての色を濡れたらtrue
	return true
}

func solve() {
	for i := 0; i < n; i++ {
		if colors[i] == 0 {
			if !dfs(i, 1) {
				println("No")
				return
			}
		}
	}

	println("Yes")
}
