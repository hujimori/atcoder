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

var (
	N   int
	S   int
	a   []int
	sum []int
)

func main() {
	sc.Split(bufio.ScanWords)
	N, S = nextInt(), nextInt()
	a = make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = nextInt()
	}
	sum = make([]int, N+1)
	// solve()
	solve2()
}

func solve2() {
	res := N + 1
	s := 0
	t := 0
	sum := 0

	for {
		for t < N && sum < S {
			sum += a[t]
			t++
		}
		if sum < S {
			break
		}

		res = min(res, t-s)
		sum -= a[s]
		s++
	}

	if res > N {
		res = 0
	}

	println(res)
}

func solve() {
	// sumを計算する
	for i := 0; i < N; i++ {
		sum[i+1] = sum[i] + a[i]
	}
	fmt.Println(sum)

	if sum[N] < S {
		// 解が存在しない
		println(0)
		return
	}

	res := N
	for s := 0; sum[s]+S <= sum[N]; s++ {
		// fmt.Printf("%d+%d <= %d\n", sum[s], S, sum[N])
		t := lowerBound(sum[s:N+1], sum[s]+S)
		res = min(res, t)
	}

	fmt.Println(res)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func lowerBound(t []int, k int) int {
	fmt.Print(t)
	fmt.Printf(" k=%d", k)
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
	fmt.Printf(" ub=%d t[%d]=%d\n", ub, ub, t[ub])
	return ub
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
