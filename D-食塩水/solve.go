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

var N, K int
var w, p []int

const INF = 110.0

func main() {
	sc.Split(bufio.ScanWords)

	N, K = nextInt(), nextInt()
	w = make([]int, N)
	p = make([]int, N)

	for i := 0; i < N; i++ {
		w[i], p[i] = nextInt(), nextInt()
	}

	lb := 0.0
	ub := INF
	for i := 0; i < 110; i++ {
		mid := (ub + lb) / 2

		if C(mid) {
			lb = mid
		} else {
			ub = mid
		}
	}

	fmt.Printf("%.12f\n", ub)
}

func C(x float64) bool {

	y := make([]float64, N)

	for i := 0; i < N; i++ {
		y[i] = float64(w[i]) * (float64(p[i]) - x)
	}

	sort.Float64s(y)
	sum := 0.0
	for i := N - 1; i > N-K-1; i-- {
		sum += y[i]
	}

	if sum >= 0 {
		return true
	}
	return false
}
