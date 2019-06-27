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

func nextLine() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)
	N, X := nextInt(), nextInt()
	x := make([]int, N+1)
	x[0] = X
	for i := 1; i < N+1; i++ {
		x[i] = nextInt()
	}
	sort.Ints(x)

	n := make([]int, N+1)
	n[0] = 0
	for i := 1; i < N+1; i++ {
		n[i] = gcd(n[i-1], abs(x[i]-x[i-1]))
	}
	fmt.Println(n[N])
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
