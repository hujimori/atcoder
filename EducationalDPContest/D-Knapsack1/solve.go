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

	dp = make([][]float64, n+100)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]float64, w+100)
	}

	for i := 0; i < n; i++ {
		for sumW := 0; sumW <= w; sumW++ {

			// i番目の品物を選ぶ
			if sumW-weight[i] >= 0 {
				chmax(&dp[i+1][sumW], dp[i][sumW-weight[i]]+float64(value[i]))
			}

			// i番目の品物を選ばない
			chmax(&dp[i+1][sumW], dp[i][sumW])

		}
	}

	// 最適値の出力
	fmt.Println(int(dp[n][w]))
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
