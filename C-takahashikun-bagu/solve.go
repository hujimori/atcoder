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
	k int
	t [][]int
)

func main() {
	sc.Split(bufio.ScanWords)

	n, k = nextInt(), nextInt()
	t = make([][]int, n)
	for i := 0; i < n; i++ {
		t[i] = make([]int, k)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < k; j++ {
			t[i][j] = nextInt()
		}
	}

	if dfs(0, 0) {
		fmt.Println("Found")
	} else {
		fmt.Println("Nothing")
	}
}

var memo [10000][10000]bool

func dfs(q, v int) bool {

	if memo[q][v] {
		return true
	}

	if q == n {
		if v == 0 {
			return true
		}
		return false
	}

	for i := 0; i < k; i++ {
		memo[q][v] = dfs(q+1, t[q][i]^v)
		if memo[q][v] {
			return true
		}
	}
	return memo[q][v]
}
