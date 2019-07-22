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

var (
	n int
	m int
	x []int
	y []int
	g [][]int
)

func main() {
	sc.Split(bufio.ScanWords)
	n, m = nextInt(), nextInt()
	x = make([]int, m)
	y = make([]int, m)
	g = make([][]int, 12)
	for i := 0; i < 12; i++ {
		g[i] = make([]int, 12)
	}
	for i := 0; i < m; i++ {
		x[i], y[i] = nextInt()-1, nextInt()-1
		g[x[i]][y[i]] = 1
		g[y[i]][x[i]] = 1
	}

	solve()
}

func solve() {

	res := 1

	for bit := 0; bit < (1 << uint(n)); bit++ {
		s := make([]int, 0)
		for i := 0; i < n; i++ {
			if bit&(1<<uint(i)) == 1<<uint(i) {
				s = append(s, i)
			}
		}

		// fmt.Println(s)

		flag := true
		for _, ss1 := range s {
			for _, ss2 := range s {
				if ss1 == ss2 {
					continue
				}
				if g[ss1][ss2] != 1 {
					flag = false
				}
			}
		}
		if flag {
			res = max(res, len(s))
		}
	}

	fmt.Println(res)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
