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

var a []int
var dp []int
var (
	n int
	m int
)

func main() {
	sc.Split(bufio.ScanWords)
	// n, m := nextInt(), nextInt()
	fmt.Scan(&n, &m)

	a = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&a[i])
	}

	dp = make([]int, n+10)
	for i := 0; i < m; i++ {
		dp[a[i]] = -1
	}
	dp[0] = 1
	mod := 1000000007
	for i := 0; i < n; i++ {
		if dp[i] == -1 {
			continue
		}
		if dp[i+1] != -1 {
			dp[i+1] += dp[i]
			dp[i+1] %= mod
		}

		if dp[i+2] != -1 {
			dp[i+2] += dp[i]
			dp[i+2] %= mod
		}

	}

	fmt.Println(dp[n])

}

// var reached = [100010]bool{}

// func dfs(i int) int {

// 	if i == n {
// 		fmt.Printf("i=%d\n", i)
// 		return 1
// 	}
// 	var res int
// 	if i+1 < n+1 && s[i+1] != -1 {
// 		// reached[i+1] = true
// 		res += dfs(i + 1)
// 	}

// 	if i+2 < n+1 && s[i+2] != -1 {
// 		res += dfs(i + 2)
// 	}

// 	return res
// }
