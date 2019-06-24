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

var n int
var w int
var weight []int
var value []int
var dp [][]float64

func main() {
	sc.Split(bufio.ScanWords)

	n, w = nextInt(), nextInt()
	weight = make([]int, n)
	value = make([]int, n)
	for i := 0; i < n; i++ {
		weight[i] = nextInt()
		value[i] = nextInt()
	}

	dp = make([][]float64, n+10)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]float64, 100100)
		for j := 0; j < 100100; j++ {
			dp[i][j] = math.Inf(1)
		}
	}

	dp[0][0] = 0

	for i := 0; i < n; i++ {
		for sumV := 0; sumV < 100100; sumV++ {

			// i番目の品物を選ぶ
			if sumV-value[i] >= 0 {
				chmin(&dp[i+1][sumV], dp[i][sumV-value[i]]+float64(weight[i]))
			}

			// i番目の品物を選ばない
			chmin(&dp[i+1][sumV], dp[i][sumV])

		}
	}

	// 最適値の出力
	res := 0
	for i := 0; i < 100100; i++ {
		if dp[n][i] <= float64(w) {
			res = i
		}
	}
	// fmt.Println(dp)
	fmt.Println(res)
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
