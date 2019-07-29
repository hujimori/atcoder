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
	ans := 0
	for i := 0; i < n; i++ {
		s[i] = nextInt()
		ans += s[i]
	}

	// dp := make([]int, n+10)
	// for i := 0; i < n+10; i++ {
	// 	dp[i] = 1000000000
	// }

	// dp[0] = 0
	// for i := 0; i < n; i++ {
	// 	for j := 1; j < n; j++ {
	// 		if dp[i] + s[i]
	// 		chmax(&dp[i+j], )
	// 	}
	// }

	// sum := 0
	// for i := 0; i < n; i++ {
	// 	if (sum+s[i])%10 == 0 {
	// 		continue
	// 	}
	// 	sum += s[i]
	// }
	// ans := 0
	// for bit := uint(0); bit < (1 << uint(n)); bit++ {
	// 	ss := []int{}

	// 	for i := uint(0); i < uint(n); i++ {
	// 		if bit&(1<<i) == 1<<i {
	// 			ss = append(ss, s[i])
	// 		}
	// 	}
	// 	sum := 0
	// 	for i := 0; i < len(ss); i++ {
	// 		sum += ss[i]
	// 	}

	// 	if sum%10 != 0 {
	// 		ans = max(ans, sum)
	// 	}
	// }

	fmt.Println(ans)

}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func chmax(a *int, b int) bool {
	if *a < b {
		*a = b
		return true
	}
	return false
}
