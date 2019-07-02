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

// 蟻本P.67 重複組み合わせ
var (
	n int
	m int
	a []int
	M int
)

func main() {
	sc.Split(bufio.ScanWords)
	n, m, M = nextInt(), nextInt(), nextInt()
	a = make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}

	solve()
}

func solve() {
	dp := make([][]int, 10000)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 10000)
	}

	// 1つも選ばない方法は常に一通り
	for i := 0; i <= n; i++ {
		dp[i][0] = 1
	}

	for i := 0; i < n; i++ {
		for j := 1; j <= m; j++ {
			if j-1-a[i] >= 0 {
				// 引き算が含まれる場合には負の数にならないように注意する
				dp[i+1][j] = (dp[i+1][j-1] + dp[i][j] - dp[i][j-1-a[i]] + M) % M
			} else {
				dp[i+1][j] = (dp[i+1][j-1] + dp[i][j]) % M
			}
		}
	}

	fmt.Println(dp[n][m])
}
