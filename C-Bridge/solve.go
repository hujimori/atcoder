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

var isVisited []bool
var g [][]bool
var n int
var m int

func main() {
	sc.Split(bufio.ScanWords)
	n, m = nextInt(), nextInt()
	g = make([][]bool, n)
	for i := 0; i < n; i++ {
		g[i] = make([]bool, n)
	}

	isVisited = make([]bool, n)

	u, v := make([]int, m), make([]int, m)
	for i := 0; i < m; i++ {
		u[i], v[i] = nextInt()-1, nextInt()-1
		g[u[i]][v[i]] = true
		g[v[i]][u[i]] = true
	}

	ans := 0
	for i := 0; i < m; i++ {
		g[u[i]][v[i]] = false
		g[v[i]][u[i]] = false

		for j := 0; j < n; j++ {
			isVisited[j] = false
		}

		dfs(0)

		bridge := false
		for j := 0; j < n; j++ {
			if !isVisited[j] {
				bridge = true
			}
		}
		if bridge {
			ans++
		}

		g[u[i]][v[i]] = true
		g[v[i]][u[i]] = true

	}

	fmt.Println(ans)
}

func dfs(v int) {
	isVisited[v] = true
	for v2 := 0; v2 < n; v2++ {
		if !g[v][v2] {
			continue
		}
		if isVisited[v2] {
			continue
		}
		dfs(v2)
	}
}
