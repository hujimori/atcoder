package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

var slice [][]int
var dp [][]float64

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)

	n := nextInt()
	slice = make([][]int, n)
	for i := 0; i < n; i++ {
		slice[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			slice[i][j] = nextInt()
		}
	}

	dp = make([][]float64, 1000000)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]float64, 3)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				if j == k {
					continue
				}
				chmax(&dp[i+1][k], dp[i][j]+float64(slice[i][k]))

			}
		}
	}

	res := 0.0
	for i := 0; i < 3; i++ {
		chmax(&res, dp[n][i])
	}

	fmt.Println(int(res))
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}

	return a
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
