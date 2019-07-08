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
	N  int
	dp []int
)

var max = 110000

func main() {
	sc.Split(bufio.ScanWords)
	N = nextInt()
	dp = make([]int, max)
	for i := 0; i < max; i++ {
		dp[i] = N
	}
	dp[0] = 0
	solve2()

}

func solve() {
	// 貰う DP --- dp[n]に遷移を集める
	for n := 1; n <= N; n++ {
		for pow6 := 1; pow6 <= n; pow6 *= 6 {
			dp[n] = min(dp[n], dp[n-pow6]+1)
		}
		for pow9 := 1; pow9 <= n; pow9 *= 9 {
			dp[n] = min(dp[n], dp[n-pow9]+1)
		}
	}

	fmt.Println(dp[N])
}

func solve2() {
	// 配る DP --- dp[n]から遷移を出していく
	for n := 0; n < N; n++ {
		for pow6 := 1; pow6+n <= N; pow6 *= 6 {
			dp[n+pow6] = min(dp[n+pow6], dp[n]+1)
		}
		for pow9 := 1; pow9+n <= N; pow9 *= 9 {
			dp[n+pow9] = min(dp[n+pow9], dp[n]+1)
		}
	}
	fmt.Println(dp[N])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
