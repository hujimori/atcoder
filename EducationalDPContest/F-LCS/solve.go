package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

var (
	s  string
	t  string
	dp [3100][3100]float64
)

func main() {
	sc.Split(bufio.ScanWords)

	ss := strings.Split(nextLine(), "")
	tt := strings.Split(nextLine(), "")

	// DPループ
	for i := 0; i < len(ss); i++ {
		for j := 0; j < len(tt); j++ {
			if ss[i] == tt[j] {
				chmax(&dp[i+1][j+1], dp[i][j]+1)
			}
			chmax(&dp[i+1][j+1], dp[i+1][j])
			chmax(&dp[i+1][j+1], dp[i][j+1])

		}
	}

	var res string

	for i, j := len(ss), len(tt); i > 0 && j > 0; {
		// (i-1, j) -> (i, j)と更新されていた場合
		if dp[i][j] == dp[i-1][j] {
			i-- // dpの遷移を遡る
		} else if dp[i][j] == dp[i][j-1] {
			// (i, j-1) -> (i, j)と更新されていた場合
			j-- // dpの遷移を遡る
		} else {
			// (i-1, j-1) -> (i, j)と更新されていた場合
			res = ss[i-1] + res
			i--
			j--
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
