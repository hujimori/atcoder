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

var colors []int // 1:黒, 0:白, -1:無色
var g [][]int
var N int
var d [2][1000000]int

func main() {
	sc.Split(bufio.ScanWords)

	N = nextInt()
	g = make([][]int, N)
	for i := 0; i < N-1; i++ {
		a, b := nextInt()-1, nextInt()-1
		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	dfs(0, -1, 0, 0)
	dfs(N-1, -1, 0, 1)

	c1 := 0
	c2 := 0
	for i := 0; i < N; i++ {
		if d[0][i] <= d[1][i] {
			c1++
		} else {
			c2++
		}
	}

	if c1 > c2 {
		fmt.Println("Fennec")
	} else {
		fmt.Println("Snuke")
	}
}

func dfs(x, y, z, w int) {

	d[w][x] = z

	for i := 0; i < len(g[x]); i++ {
		if g[x][i] == y {
			continue
		}

		dfs(g[x][i], x, z+1, w)

	}
}
