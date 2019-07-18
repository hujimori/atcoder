package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/go-algorithms/binarysearch"
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

	n, m := nextInt(), nextInt()
	x, y := nextInt(), nextInt()
	a := make([]int, n)
	b := make([]int, m)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	for i := 0; i < m; i++ {
		b[i] = nextInt()
	}

	ans := 0
	sum := 0
	av := 0
	bv := 0
	for ans = 0; ; ans++ {

		if av == len(a) || bv == len(b) {
			break
		}

		bv = binarysearch.LowerBound(b, sum+x)
		if bv < len(b) {
			sum += x
			sum += abs(b[bv] - sum)
		}

		av = binarysearch.LowerBound(a, sum+y)
		if av < len(a) {
			sum += y
			sum += abs(a[av] - sum)
		}

	}

	fmt.Println(ans)
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
