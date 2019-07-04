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

var (
	n int
	k int
	w []int
	v []int

	y []float64 // v - x*w
)

func main() {
	sc.Split(bufio.ScanWords)
	n, k = nextInt(), nextInt()
	w = make([]int, n)
	v = make([]int, n)
	for i := 0; i < n; i++ {
		w[i], v[i] = nextInt(), nextInt()
	}
	sort.Ints(w)
	sort.Ints(v)

	y = make([]float64, n)

	solve()
}

func C(x float64) bool {
	for i := 0; i < n; i++ {
		y[i] = float64(v[i]) - x*float64(w[i])
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(y)))

	sum := 0.0
	for i := 0; i < k; i++ {
		sum += y[i]
	}

	return sum >= 0.0
}

func solve() {
	lb := 0.0
	ub := 10000000000000.0
	for i := 0; i < 100; i++ {
		mid := (ub + lb) / 2
		fmt.Printf("mid=%f\n", mid)
		if C(mid) {
			lb = mid
		} else {
			ub = mid
		}
	}
	fmt.Printf("%.2f\n", ub)
}
