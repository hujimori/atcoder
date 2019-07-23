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
	c := make([]int, n)
	for i := 0; i < n; i++ {
		c[i] = nextInt()
	}

	inf := 100000001000000
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = inf
	}

	for i := 0; i < n; i++ {

		key := LowerBound(dp, c[i])
		dp[key] = c[i]
		// fmt.Printf("key=%d i=%d\n", key, i)
		// fmt.Println(dp[0 : n-1])

	}

	ans := 0
	for i := 0; i < n; i++ {
		if dp[i] == inf {
			ans++
		}
	}

	fmt.Println(ans)
}

func LowerBound(t []int, k int) int {
	lb := -1
	ub := len(t)
	for ub-lb > 1 {
		mid := (lb + ub) / 2

		if t[mid] >= k {
			ub = mid
		} else {
			lb = mid
		}
	}
	return ub
}
