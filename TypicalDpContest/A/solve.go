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

const MAXN = 110
const MAXC = 110

var (
	dp [10001]int
	p  []int
	n  int
)

func main() {
	sc.Split(bufio.ScanWords)
	n = nextInt()
	p = make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = nextInt()
	}

	dp[0] = 1

	for i := 0; i < n; i++ {
		for j := 10000; j >= 0; j-- {
			if dp[j] == 1 {
				dp[j+p[i]] = 1
			}
		}
	}

	ans := 0
	for i := 0; i < len(dp); i++ {
		ans += dp[i]
	}

	fmt.Println(ans)

	// for i := 0; i < len(dp); i++ {
	// 	// dp[i] = math.Inf(1)
	// }

	// dp[0] = 0

	// for i := 0; i < n; i++ {
	// 	for count := 0; count <= n; count++ {

	// 		// i番目の問題を解く場合
	// 		if count-1 >= 0 {
	// 			chmax(&dp[i+1][count], dp[i][count-1]+float64(p[i]))
	// 		}
	// 		// i番目の問題を解かない場合
	// 		chmax(&dp[i+1][count], dp[i][count])

	// 	}
	// }

	// for i := 0; i < MAXN; i++ {
	// 	for j := 0; j < MAXC; j++ {
	// 		if dp[i][j] != 0 {
	// 			fmt.pr
	// 			fmt.Println(dp[i][j])
	// 		}
	// 	}
	// }
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
