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

var (
	n      int
	w      int
	weight []int
	value  []int
)

const MAXN = 110
const MAXV = 100100

var dp [MAXN][MAXV]float64

func main() {
	sc.Split(bufio.ScanWords)

	n, w = nextInt(), nextInt()
	weight = make([]int, n)
	value = make([]int, n)
	for i := 0; i < n; i++ {
		weight[i] = nextInt()
		value[i] = nextInt()
	}

	for i := 0; i < MAXN; i++ {
		for j := 0; j < MAXV; j++ {
			dp[i][j] = math.Inf(1)
		}
	}

	dp[0][0] = 0

	for i := 0; i < n; i++ {
		for sumv := 0; sumv < MAXV; sumv++ {
			// i番目の品物を選ぶ
			if sumv-value[i] >= 0 {
				chmin(&dp[i+1][sumv], (dp[i][sumv-value[i]] + float64(weight[i])))
			}
			// i番目の品物を選ばない
			chmin(&dp[i+1][sumv], dp[i][sumv])

		}
	}

	// 最適値の出力
	res := 0
	for sumv := 0; sumv < MAXV; sumv++ {
		if dp[n][sumv] <= float64(w) {
			res = int(sumv)
		}
	}

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
