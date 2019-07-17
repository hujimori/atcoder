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

var w int
var n int
var k int
var a []int
var b []int

var dp [55][55][10010]int

func main() {
	sc.Split(bufio.ScanWords)

	w = nextInt()
	n, k = nextInt(), nextInt()

	a = make([]int, n)
	b = make([]int, n)
	for i := 0; i < n; i++ {
		a[i], b[i] = nextInt(), nextInt()
	}

	for i := 0; i < 55; i++ {
		for j := 0; j < 55; j++ {
			for k := 0; k < 10010; k++ {
				dp[i][j][k] = -1
			}
		}
	}

	dp[0][0][0] = 0

	for i := 0; i < n; i++ {
		for j := 0; j <= k; j++ {
			for sumw := 0; sumw <= w; sumw++ {
				if sumw+a[i] <= w {
					// 選べる場合
					dp[i+1][j+1][sumw+a[i]] = max(dp[i+1][j+1][sumw+a[i]], dp[i][j][sumw]+b[i])
				}
				// 選べない場合
				dp[i+1][j][sumw] = max(dp[i][j][sumw], dp[i+1][j][sumw])
			}
		}
	}

	ans := -1
	for j := 0; j <= k; j++ {
		for sumw := 0; sumw < 10010; sumw++ {
			ans = max(ans, dp[n][j][sumw])
		}
	}

	fmt.Println(ans)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
