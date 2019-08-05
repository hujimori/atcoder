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

func main() {
	sc.Split(bufio.ScanWords)
	n, x := nextInt(), nextInt()

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}

	ans := 0
	for i := 0; i < n-1; i++ {

		delta := max(0, a[i]+a[i+1]-x)
		a[i+1] -= delta

		if a[i+1] < 0 {
			a[i] += a[i+1]
			a[i+1] = 0
		}

		ans += delta

	}

	fmt.Println(ans)
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
