package main

import (
	"bufio"
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

var (
	n int
	m int
	x []int
)

func main() {
	sc.Split(bufio.ScanWords)
	n, m = nextInt(), nextInt()
	x = make([]int, n)
	for i := 0; i < n; i++ {
		x[i] = nextInt()
	}

	sort.Ints(x)

	solve()
}

// 条件を満たすか判定
func C(d int) bool {
	last := 0

	for i := 1; i < m; i++ {
		crt := last + 1
		for crt < n && x[crt]-x[last] < d {
			crt++
		}
		if crt == n {
			return false
		}
		last = crt
	}
	return true
}

func solve() {
	lb := 0
	ub := 1000000000000

	for ub-lb > 1 {
		mid := (lb + ub) / 2
		if C(mid) {
			lb = mid
		} else {
			ub = mid
		}
	}
	println(lb)
}
