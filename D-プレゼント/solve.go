package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

type Input struct {
	first  int
	second int
}

type Inputs []Input

func (in Inputs) Len() int {
	return len(in)
}

func (in Inputs) Swap(i, j int) {
	in[i], in[j] = in[j], in[i]
}

func (in Inputs) Less(i, j int) bool {
	if in[i].second < in[j].second {
		return true

	} else if in[i].second == in[j].second {
		if in[i].first > in[j].first {
			return true
		}
	}

	return false
}

func main() {
	sc.Split(bufio.ScanWords)
	n := nextInt()
	in := make([]Input, n)
	for i := 0; i < n; i++ {
		in[i].first = nextInt()
		in[i].second = nextInt()
	}

	sort.Sort(Inputs(in))
	// fmt.Println(in)

	inf := 100000001000000
	dp := make([]int, n+10)
	for i := 0; i < n+10; i++ {
		dp[i] = inf
	}

	for i := 0; i < n; i++ {

		key := LowerBound(dp, in[i].first)
		dp[key] = in[i].first

		// fmt.Println(dp[0 : n-1])

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
