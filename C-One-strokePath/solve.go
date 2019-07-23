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

var g [][]bool
var n, m int

func main() {
	sc.Split(bufio.ScanWords)
	n, m = nextInt(), nextInt()
	g = make([][]bool, n)
	for i := 0; i < n; i++ {
		g[i] = make([]bool, n)
	}
	for i := 0; i < m; i++ {
		a, b := nextInt()-1, nextInt()-1
		g[a][b] = true
		g[b][a] = true
	}

	visted := make([]bool, n)
	visted[0] = true
	fmt.Println(dfs(0, visted))
}

func dfs(v int, visted []bool) int {
	can := true
	for i := 0; i < n; i++ {
		if !visted[i] {
			can = false
		}
	}

	if can {
		return 1
	}

	ret := 0

	for i := 0; i < n; i++ {
		if !g[v][i] {
			continue
		}

		if visted[i] {
			continue
		}

		visted[i] = true
		ret += dfs(i, visted)
		visted[i] = false
	}

	return ret
}
