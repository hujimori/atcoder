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
	N int
	A int
	B int
	C int
)

var l []int

func main() {
	sc.Split(bufio.ScanWords)
	N, A, B, C = nextInt(), nextInt(), nextInt(), nextInt()
	// fmt.Scan(&N, &A, &B, &C)
	l = make([]int, N)
	for i := 0; i < N; i++ {
		l[i] = nextInt()
	}

	fmt.Println(dfs(0, 0, 0, 0))

}

func dfs(i, a, b, c int) int {
	// fmt.Println(l)
	if i == N {
		if a == 0 || b == 0 || c == 0 {
			return 1000000000
		}
		return abs(a-A) + abs(b-B) + abs(c-C)
	}

	res := dfs(i+1, a, b, c)

	if a != 0 {
		chmin(&res, dfs(i+1, a+l[i], b, c)+10)
	} else {
		chmin(&res, dfs(i+1, a+l[i], b, c))
	}

	if b != 0 {
		chmin(&res, dfs(i+1, a, b+l[i], c)+10)
	} else {
		chmin(&res, dfs(i+1, a, b+l[i], c))
	}

	if c != 0 {
		chmin(&res, dfs(i+1, a, b, c+l[i])+10)
	} else {
		chmin(&res, dfs(i+1, a, b, c+l[i]))
	}
	return res
}

func chmin(a *int, b int) bool {
	if *a > b {
		*a = b
		return true
	}
	return false
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}

	return a
}
