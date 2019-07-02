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

// 蟻本P.66 分割数
var (
	n int
	m int
	M int
)

func main() {
	sc.Split(bufio.ScanWords)
	n, m, M = nextInt(), nextInt(), nextInt()
	solve()
}

func solve() {
	dp := make([][]int, 1000)
	for i := 0; i < 1000; i++ {
		dp[i] = make([]int, 1000)
	}

	dp[0][0] = 1
	for i := 1; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if j-i >= 0 {
				dp[i][j] = (dp[i-1][j] + dp[i][j-i]) % M
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	fmt.Println(dp[m][n])
}
