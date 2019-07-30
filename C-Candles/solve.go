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

	n, k := nextInt(), nextInt()

	x := make([]int, n)
	for i := 0; i < n; i++ {
		x[i] = nextInt()
	}

	right := 0
	left := 0
	ans := 1 << 60
	for i := 0; i+k-1 < n; i++ {
		left = x[i]
		right = x[i+k-1]
		ans = min(ans, min(abs(left), abs(right))+right-left)
	}

	fmt.Println(ans)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}
