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

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

var (
	N  int
	W  int
	v  []int64
	w  []int64
	dp [][]int64
)

var INF = 10000000100

func main() {
	sc.Split(bufio.ScanWords)
	N, W = nextInt(), nextInt()
	v = make([]int64, N)
	w = make([]int64, N)
	for i := 0; i < N; i++ {
		v[i], w[i] = int64(nextInt()), int64(nextInt())
	}

	dp = make([][]int64, N+2)
	for i := 0; i < N+2; i++ {
		dp[i] = make([]int64, INF)
	}

	for i := 0; i < INF; i++ {
		dp[0][i] = int64(INF)
	}

	dp[0][0] = 0
	for i := 0; i < N; i++ {
		for j := 0; j < INF; j++ {
			if int64(j) < v[i] {
				dp[i+1][j] = dp[i][j]
			} else {
				dp[i+1][j] = min(dp[i][j], dp[i][j-int(v[i])]+w[i])
			}
		}
	}

	res := 0
	for i := 0; i < INF; i++ {
		if dp[N][i] <= int64(W) {
			res = i
		}
	}

	fmt.Println(res)

}

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}
