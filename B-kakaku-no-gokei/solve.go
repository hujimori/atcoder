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
	n int
	x int
	a []int
)

func main() {
	sc.Split(bufio.ScanWords)

	n, x = nextInt(), nextInt()
	a = make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt()
	}
	solve()
}

func solve() {
	var slice []int
	for i := uint(0); i < uint(n); i++ {
		if x&(1<<i) == 1<<i {
			slice = append(slice, a[i])
		}
	}

	ans := 0
	for _, s := range slice {
		ans += s
	}
	fmt.Println(ans)
}
