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

func main() {
	sc.Split(bufio.ScanWords)

	n := nextInt()
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = nextInt()
	}

	var dp [101][10101]bool
	dp[0][0] = true

	for i := 0; i < n; i++ {
		for j := 0; j < 10001; j++ {
			if dp[i][j] {

				dp[i+1][j] = true
				dp[i+1][j+s[i]] = true
			}
		}
	}

	ans := 0
	for j := 0; j < 10001; j++ {
		if j%10 != 0 && dp[n][j] {
			ans = max(ans, j)
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
