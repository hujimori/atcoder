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

const mod = 1000000007

var dp [110000]int
var sum [110000]int

func main() {
	sc.Split(bufio.ScanWords)
	n, m := nextInt(), nextInt()
	f := make([]int, n)

	for i := 0; i < n; i++ {
		f[i] = nextInt()
		f[i]--
	}

	fnum := make([]int, m) // 区間に種類iが何個あるか
	L := make([]int, n+1)  // 各iに対する尺取法の区間
	left := 0
	for right := 0; right < n; right++ {
		fnum[f[right]]++
		for left < right && fnum[f[right]] > 1 {
			fnum[f[left]]--
			left++
		}
		L[right+1] = left
	}

	// 累積和で高速化したDP
	dp[0] = 1
	sum[0] = 0
	sum[1] = 1
	for i := 1; i <= n; i++ {
		dp[i] = (sum[i] - sum[L[i]] + mod) % mod // DP
		sum[i+1] = (sum[i] + dp[i]) % mod        // 累積
	}

	fmt.Println(dp[n])

}
