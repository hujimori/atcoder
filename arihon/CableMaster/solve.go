package main

import (
	"bufio"
	"fmt"
	"math"
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

var (
	n int
	k int
	l []float64
)

func main() {
	// sc.Split(bufio.ScanWords)
	// n, k = nextInt(), nextInt()
	fmt.Scan(&n)
	fmt.Scan(&k)
	l = make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&l[i])
	}
	fmt.Println(l)

	solve()
}

func C(x float64) bool {
	num := 0
	for i := 0; i < n; i++ {
		num += int(l[i] / x)
	}
	return num >= k
}

func solve() {
	// 解の存在範囲を初期化
	lb := 0.0
	ub := 1000000000000000000.0
	// ub := math.MaxFloat64
	// 解の存在が十分なくなるまで繰り返す
	for i := 0; i < 100; i++ {
		mid := (lb + ub) / 2.0
		// fmt.Printf("mid=%f\n", mid)
		if C(mid) {
			lb = mid
		} else {
			ub = mid
		}
	}

	fmt.Printf("%.2f\n", math.Floor(ub*100)/100)
}
