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
	box := make([]int, n)
	for i := 0; i < n; i++ {
		box[i] = nextInt()
	}

	inf := 1000000001000
	dp := make([]int, n+10)
	for i := 0; i < n+10; i++ {
		dp[i] = inf
	}

	for i := 0; i < n; i++ {
		key := LowerBound(dp, box[i])
		dp[key] = box[i]
	}

	ans := 0
	for i := 0; i < n+10; i++ {
		if dp[i] < inf {
			ans = i + 1
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
