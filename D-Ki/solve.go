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

var g [][]int
var val []int

func dfs(p, c, sum int) {
	val[c] += sum

	for _, v := range g[c] {
		if v != p {
			dfs(c, v, val[c])
		}
	}
}

func main() {
	sc.Split(bufio.ScanWords)

	n, q := nextInt(), nextInt()
	g = make([][]int, n)
	val = make([]int, n)
	for i := 0; i < n-1; i++ {
		a, b := nextInt(), nextInt()
		a--
		b--

		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	for i := 0; i < q; i++ {
		p, x := nextInt(), nextInt()
		p--
		val[p] += x
	}

	dfs(-1, 0, 0)

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", val[i])
	}

	fmt.Println()
}
