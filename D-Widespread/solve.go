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

const INF = 100001000000

var h []int
var N, A, B int

func main() {
	sc.Split(bufio.ScanWords)

	N, A, B = nextInt(), nextInt(), nextInt()
	h = make([]int, N)
	for i := 0; i < N; i++ {
		h[i] = nextInt()
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

func C(t int) bool {

	cnt := 0

	for _, hh := range h {
		if hh > B*t {
			cnt += (hh - B*t) / (A - B)
			if (hh-B*t)%(A-B) != 0 {
				cnt++
			}
		}
	}

	if t >= cnt {
		return true
	}
	return false

}
