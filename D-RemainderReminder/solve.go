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
	sum := 0

	for b := 1; b <= n; b++ {
		sum += max(0, b-k)*(n/b) + max(0, n%b+1-k)
	}
	if k == 0 {
		sum = n * n
	}

	fmt.Println(sum)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
