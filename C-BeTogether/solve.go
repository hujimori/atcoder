package main

import (
	"bufio"
	"fmt"
	"math"
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

func main() {
	sc.Split(bufio.ScanWords)

	n := nextInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}

	sort.Ints(a)
	// s := make([]int, n)
	// s[0] = 0
	// for i := 1; i < n; i++ {
	// 	s[i] = gcd(s[i-1], abs(a[i]-a[i-1]))
	// }
	start := a[0]
	end := a[n-1]

	ans := math.Inf(1)

	for s := start; s <= end; s++ {
		cost := 0
		for i := 0; i < n; i++ {
			cost += (a[i] - s) * (a[i] - s)
		}
		if ans > float64(cost) {
			ans = float64(cost)
		}
	}

	fmt.Println(int(ans))
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
