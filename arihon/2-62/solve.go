package main

import (
	"bufio"
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

// 蟻本P.62 個数制限付き部分和問題
var n int
var a []int
var m []int
var k int

// var dp [][]bool

func main() {
	sc.Split(bufio.ScanWords)
	n, k = nextInt(), nextInt()
	a = make([]int, n)
	m = make([]int, n)
	for i := 0; i < n; i++ {
		a[i], m[i] = nextInt(), nextInt()
	}
	// solve()
	solve2()

}

func solve() {
	var dp [][]bool
	dp = make([][]bool, 100000)
	for i := 0; i < 100000; i++ {
		dp[i] = make([]bool, 100000)
	}
	dp[0][0] = true

	for i := 0; i < n; i++ {
		for j := 0; j <= k; j++ {
			for k := 0; k <= m[i] && k*a[i] <= j; k++ {
				dp[i+1][j] = dp[i+1][j] || dp[i][j-k*a[i]]
			}
		}
	}
	if dp[n][k] {
		println("Yes")
	} else {
		println("No")
	}
}

func solve2() {
	var dp []int
	dp = make([]int, 100000)
	for i := 0; i < 100000; i++ {
		dp[i] = -1
	}
	dp[0] = 0
	for i := 0; i < n; i++ {
		for j := 0; j <= k; j++ {
			if dp[j] >= 0 {
				dp[j] = m[i]
			} else if j < a[i] || dp[j-a[i]] <= 0 {
				dp[j] = -1
			} else {
				dp[j] = dp[j-a[i]] - 1
			}
		}
	}

	if dp[k] >= 0 {
		println("Yes")
	} else {
		println("No")
	}
}

func chmin(a *float64, b float64) bool {
	if *a > b {
		*a = b
		return true
	}
	return false
}

func chmax(a *float64, b float64) bool {
	if *a < b {
		*a = b
		return true
	}
	return false
}
