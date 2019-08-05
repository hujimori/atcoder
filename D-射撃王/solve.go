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

var N int
var H, S []int

const INF = 1000010000000000

func main() {
	sc.Split(bufio.ScanWords)

	N = nextInt()
	H = make([]int, N)
	S = make([]int, N)
	for i := 0; i < N; i++ {
		H[i], S[i] = nextInt(), nextInt()
	}

	lb := 0
	ub := INF

	for ub-lb > 1 {
		mid := (ub + lb) / 2
		if C(mid) {
			ub = mid
		} else {
			lb = mid
		}
	}

	if !C(ub) {
		ub = lb
	}

	fmt.Println(ub)
}

func C(x int) bool {
	t := make([]int, N)
	for i := 0; i < N; i++ {
		if (x - H[i]) < 0 {
			return false
		}
		// t[i] = min(1111111100, (x-H[i])/S[i])
		t[i] = (x - H[i]) / S[i]
	}

	sort.Ints(t)
	for i := 0; i < N; i++ {
		if t[i] < i {
			return false
		}
	}
	return true

}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
