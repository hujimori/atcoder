package main

import (
	"bufio"
	"fmt"
	"math"
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

// 蟻本P.64 最長増加部分列問題
var n int
var a []int

func main() {
	sc.Split(bufio.ScanWords)
	n = nextInt()
	a = make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}

	solve()

}

func solve() {
	dp := make([]int, 1000)
	res := 0.0
	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if a[j] < a[i] {
				dp[i] = int(math.Max(float64(dp[i]), float64(dp[j]+1)))
			}
			chmax(&res, float64(dp[i]))
		}
	}

	fmt.Println(int(res))
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
